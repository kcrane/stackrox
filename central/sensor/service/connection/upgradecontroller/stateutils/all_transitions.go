package stateutils

import (
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/sensorupgrader"
)

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func statePtr(state storage.UpgradeProgress_UpgradeState) *storage.UpgradeProgress_UpgradeState {
	return &state
}

var (
	// These define all the valid transitions we handle.
	// Note that the first match wins, so order these transitions with that in mind.
	allTransitions = []transitioner{
		// If the upgrade is in a terminal state, just tell it to clean up.
		{
			currentStateMatch: anyStateFrom(TerminalStates.AsSlice()...),

			noStateChange:     true,
			workflowToExecute: sensorupgrader.CleanupWorkflow,
		},

		// The following transitions handle the situation right after the upgrader comes up.
		// (Indicated by an empty string for the workflow.)
		// Note that the upgrader might restart at any time.
		// So we MUST handle all non-terminal states through the below transitions.

		{
			workflowMatch: stringPtr(""),
			currentStateMatch: anyStateFrom(
				storage.UpgradeProgress_UNSET, // This should basically never happen, but being defensive can't hurt.

				// These two states would be a little early to hear from the upgrader, but still possible in case
				// the upgrader happens to reach out before sensor for whatever reason.
				storage.UpgradeProgress_UPGRADE_TRIGGER_SENT,
				storage.UpgradeProgress_UPGRADER_LAUNCHING,

				storage.UpgradeProgress_UPGRADER_LAUNCHED, // This is the stage where we normally expect to hear from the upgrader.

				// Seeing the below states likely means the upgrader restarted part way through the process. However, we haven't heard
				// from the sensor yet (else we'd say upgrade complete), and the upgrader is idempotent, so tell it to roll-forward anyway.
				storage.UpgradeProgress_PRE_FLIGHT_CHECKS_COMPLETE,
				storage.UpgradeProgress_UPGRADE_OPERATIONS_DONE,
			),

			workflowToExecute: sensorupgrader.RollForwardWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADER_LAUNCHED),
		},
		{
			// Upgrader restarted in the middle of rolling back. Tell it to keep rolling back.
			workflowMatch:     stringPtr(""),
			currentStateMatch: anyStateFrom(storage.UpgradeProgress_UPGRADE_ERROR_ROLLING_BACK),

			workflowToExecute: sensorupgrader.RollBackWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADE_ERROR_ROLLING_BACK),
		},

		// The following are roll-forward transitions.
		// Note that we don't check the starting state here (we know it's not terminal since that was checked above,
		// and the end state only depends on the upgrader action).
		{
			workflowMatch:    stringPtr(sensorupgrader.RollForwardWorkflow),
			stageMatch:       anyStageFrom(rollForwardStagesBeforePreFlight...),
			errOccurredMatch: boolPtr(false),

			workflowToExecute: sensorupgrader.RollForwardWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADER_LAUNCHED),
		},
		// An error occurred before we could even do pre-flight checks!
		// Mark it as a fatal error, and tell the upgrader to clean up.
		{
			workflowMatch:    stringPtr(sensorupgrader.RollForwardWorkflow),
			stageMatch:       anyStageFrom(rollForwardStagesBeforePreFlight...),
			errOccurredMatch: boolPtr(true),

			workflowToExecute: sensorupgrader.CleanupWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADE_INITIALIZATION_ERROR),
			updateDetail:      true,
		},
		// Yay, passed pre-flight checks!
		{
			workflowMatch:    stringPtr(sensorupgrader.RollForwardWorkflow),
			stageMatch:       anyStageFrom(sensorupgrader.PreflightStage),
			errOccurredMatch: boolPtr(false),

			workflowToExecute: sensorupgrader.RollForwardWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_PRE_FLIGHT_CHECKS_COMPLETE),
		},
		// Oh no, pre-flight checks failed!
		{
			workflowMatch:    stringPtr(sensorupgrader.RollForwardWorkflow),
			stageMatch:       anyStageFrom(sensorupgrader.PreflightStage),
			errOccurredMatch: boolPtr(true),

			workflowToExecute: sensorupgrader.CleanupWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_PRE_FLIGHT_CHECKS_FAILED),
			updateDetail:      true,
		},
		// Ooh yeah, upgrade done from the PoV of the upgrader!
		{
			workflowMatch:    stringPtr(sensorupgrader.RollForwardWorkflow),
			stageMatch:       anyStageFrom(sensorupgrader.ExecuteStage),
			errOccurredMatch: boolPtr(false),

			// Tell the upgrader to stay in the roll-forward workflow, and keep polling until
			// we ask it to clean up (after we hear from the sensor).
			workflowToExecute: sensorupgrader.RollForwardWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADE_OPERATIONS_DONE),
		},
		// Oh no, upgrade operations failed. :( Tell the upgrader to roll back.
		{
			workflowMatch:    stringPtr(sensorupgrader.RollForwardWorkflow),
			stageMatch:       anyStageFrom(sensorupgrader.ExecuteStage),
			errOccurredMatch: boolPtr(true),

			workflowToExecute: sensorupgrader.RollBackWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADE_ERROR_ROLLING_BACK),
			updateDetail:      true,
		},

		// The following are roll-back transitions.

		// Rollback still in progress.
		{
			workflowMatch:    stringPtr(sensorupgrader.RollBackWorkflow),
			stageMatch:       anyStageFrom(sensorupgrader.SnapshotForRollbackStage, sensorupgrader.GenerateRollbackPlanStage, sensorupgrader.PreflightNoFailStage),
			errOccurredMatch: boolPtr(false),

			workflowToExecute: sensorupgrader.RollBackWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADE_ERROR_ROLLING_BACK),
		},
		// Rollback done, now clean up.
		{
			workflowMatch:    stringPtr(sensorupgrader.RollBackWorkflow),
			stageMatch:       anyStageFrom(sensorupgrader.ExecuteStage),
			errOccurredMatch: boolPtr(false),

			workflowToExecute: sensorupgrader.CleanupWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADE_ERROR_ROLLED_BACK),
		},
		// Any error when rolling back => rollback failed. Not much we can do at this point. :(
		{
			workflowMatch:    stringPtr(sensorupgrader.RollBackWorkflow),
			errOccurredMatch: boolPtr(true),

			// Upgrader might as well clean up.
			workflowToExecute: sensorupgrader.CleanupWorkflow,
			nextState:         statePtr(storage.UpgradeProgress_UPGRADE_ERROR_ROLLBACK_FAILED),
			updateDetail:      true,
		},

		// No need to define transitions explicitly for clean up since the upgrader
		// should only be cleaning up on terminal states.
	}
)
