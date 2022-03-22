// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

var (
	log = logging.LoggerForModule()
)

const (
	baseTable  = "cluster_health_status"
	countStmt  = "SELECT COUNT(*) FROM cluster_health_status"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM cluster_health_status WHERE Id = $1)"

	getStmt     = "SELECT serialized FROM cluster_health_status WHERE Id = $1"
	deleteStmt  = "DELETE FROM cluster_health_status WHERE Id = $1"
	walkStmt    = "SELECT serialized FROM cluster_health_status"
	getIDsStmt  = "SELECT Id FROM cluster_health_status"
	getManyStmt = "SELECT serialized FROM cluster_health_status WHERE Id = ANY($1::text[])"

	deleteManyStmt = "DELETE FROM cluster_health_status WHERE Id = ANY($1::text[])"
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.ClusterHealthStatus)(nil)), baseTable)
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*storage.ClusterHealthStatus, bool, error)
	Upsert(ctx context.Context, obj *storage.ClusterHealthStatus) error
	UpsertMany(ctx context.Context, objs []*storage.ClusterHealthStatus) error
	Delete(ctx context.Context, id string) error
	GetIDs(ctx context.Context) ([]string, error)
	GetMany(ctx context.Context, ids []string) ([]*storage.ClusterHealthStatus, []int, error)
	DeleteMany(ctx context.Context, ids []string) error

	Walk(ctx context.Context, fn func(obj *storage.ClusterHealthStatus) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableClusterHealthStatus(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists cluster_health_status (
    Id varchar,
    CollectorHealthInfo_Version varchar,
    CollectorHealthInfo_TotalDesiredPods integer,
    CollectorHealthInfo_TotalReadyPods integer,
    CollectorHealthInfo_TotalRegisteredNodes integer,
    CollectorHealthInfo_StatusErrors text[],
    AdmissionControlHealthInfo_TotalDesiredPods integer,
    AdmissionControlHealthInfo_TotalReadyPods integer,
    AdmissionControlHealthInfo_StatusErrors text[],
    SensorHealthStatus integer,
    CollectorHealthStatus integer,
    OverallHealthStatus integer,
    AdmissionControlHealthStatus integer,
    LastContact timestamp,
    HealthInfoComplete bool,
    serialized bytea,
    PRIMARY KEY(Id)
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoClusterHealthStatus(ctx context.Context, tx pgx.Tx, obj *storage.ClusterHealthStatus) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetCollectorHealthInfo().GetVersion(),
		obj.GetCollectorHealthInfo().GetTotalDesiredPods(),
		obj.GetCollectorHealthInfo().GetTotalReadyPods(),
		obj.GetCollectorHealthInfo().GetTotalRegisteredNodes(),
		obj.GetCollectorHealthInfo().GetStatusErrors(),
		obj.GetAdmissionControlHealthInfo().GetTotalDesiredPods(),
		obj.GetAdmissionControlHealthInfo().GetTotalReadyPods(),
		obj.GetAdmissionControlHealthInfo().GetStatusErrors(),
		obj.GetSensorHealthStatus(),
		obj.GetCollectorHealthStatus(),
		obj.GetOverallHealthStatus(),
		obj.GetAdmissionControlHealthStatus(),
		pgutils.NilOrStringTimestamp(obj.GetLastContact()),
		obj.GetHealthInfoComplete(),
		serialized,
	}

	finalStr := "INSERT INTO cluster_health_status (Id, CollectorHealthInfo_Version, CollectorHealthInfo_TotalDesiredPods, CollectorHealthInfo_TotalReadyPods, CollectorHealthInfo_TotalRegisteredNodes, CollectorHealthInfo_StatusErrors, AdmissionControlHealthInfo_TotalDesiredPods, AdmissionControlHealthInfo_TotalReadyPods, AdmissionControlHealthInfo_StatusErrors, SensorHealthStatus, CollectorHealthStatus, OverallHealthStatus, AdmissionControlHealthStatus, LastContact, HealthInfoComplete, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, CollectorHealthInfo_Version = EXCLUDED.CollectorHealthInfo_Version, CollectorHealthInfo_TotalDesiredPods = EXCLUDED.CollectorHealthInfo_TotalDesiredPods, CollectorHealthInfo_TotalReadyPods = EXCLUDED.CollectorHealthInfo_TotalReadyPods, CollectorHealthInfo_TotalRegisteredNodes = EXCLUDED.CollectorHealthInfo_TotalRegisteredNodes, CollectorHealthInfo_StatusErrors = EXCLUDED.CollectorHealthInfo_StatusErrors, AdmissionControlHealthInfo_TotalDesiredPods = EXCLUDED.AdmissionControlHealthInfo_TotalDesiredPods, AdmissionControlHealthInfo_TotalReadyPods = EXCLUDED.AdmissionControlHealthInfo_TotalReadyPods, AdmissionControlHealthInfo_StatusErrors = EXCLUDED.AdmissionControlHealthInfo_StatusErrors, SensorHealthStatus = EXCLUDED.SensorHealthStatus, CollectorHealthStatus = EXCLUDED.CollectorHealthStatus, OverallHealthStatus = EXCLUDED.OverallHealthStatus, AdmissionControlHealthStatus = EXCLUDED.AdmissionControlHealthStatus, LastContact = EXCLUDED.LastContact, HealthInfoComplete = EXCLUDED.HealthInfoComplete, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	createTableClusterHealthStatus(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.ClusterHealthStatus) error {
	conn, release := s.acquireConn(ctx, ops.Get, "ClusterHealthStatus")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoClusterHealthStatus(ctx, tx, obj); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.ClusterHealthStatus) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "ClusterHealthStatus")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.ClusterHealthStatus) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "ClusterHealthStatus")

	return s.upsert(ctx, objs...)
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "ClusterHealthStatus")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "ClusterHealthStatus")

	row := s.db.QueryRow(ctx, existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string) (*storage.ClusterHealthStatus, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "ClusterHealthStatus")

	conn, release := s.acquireConn(ctx, ops.Get, "ClusterHealthStatus")
	defer release()

	row := conn.QueryRow(ctx, getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.ClusterHealthStatus
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(ctx context.Context, id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "ClusterHealthStatus")

	conn, release := s.acquireConn(ctx, ops.Remove, "ClusterHealthStatus")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs(ctx context.Context) ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "storage.ClusterHealthStatusIDs")

	rows, err := s.db.Query(ctx, getIDsStmt)
	if err != nil {
		return nil, pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice
func (s *storeImpl) GetMany(ctx context.Context, ids []string) ([]*storage.ClusterHealthStatus, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "ClusterHealthStatus")

	conn, release := s.acquireConn(ctx, ops.GetMany, "ClusterHealthStatus")
	defer release()

	rows, err := conn.Query(ctx, getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	resultsByID := make(map[string]*storage.ClusterHealthStatus)
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := &storage.ClusterHealthStatus{}
		if err := proto.Unmarshal(data, msg); err != nil {
			return nil, nil, err
		}
		resultsByID[msg.GetId()] = msg
	}
	missingIndices := make([]int, 0, len(ids)-len(resultsByID))
	// It is important that the elems are populated in the same order as the input ids
	// slice, since some calling code relies on that to maintain order.
	elems := make([]*storage.ClusterHealthStatus, 0, len(resultsByID))
	for i, id := range ids {
		if result, ok := resultsByID[id]; !ok {
			missingIndices = append(missingIndices, i)
		} else {
			elems = append(elems, result)
		}
	}
	return elems, missingIndices, nil
}

// Delete removes the specified IDs from the store
func (s *storeImpl) DeleteMany(ctx context.Context, ids []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "ClusterHealthStatus")

	conn, release := s.acquireConn(ctx, ops.RemoveMany, "ClusterHealthStatus")
	defer release()
	if _, err := conn.Exec(ctx, deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.ClusterHealthStatus) error) error {
	rows, err := s.db.Query(ctx, walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.ClusterHealthStatus
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableClusterHealthStatus(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS cluster_health_status CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableClusterHealthStatus(ctx, db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}
