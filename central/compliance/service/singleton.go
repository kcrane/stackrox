package service

import (
	"sync"

	"github.com/stackrox/rox/central/cluster/datastore"
	"github.com/stackrox/rox/central/compliance/standards"
	"github.com/stackrox/rox/central/compliance/store"
)

var (
	serviceInstance     Service
	serviceInstanceInit sync.Once
)

// Singleton returns the singleton instance of the compliance service.
func Singleton() Service {
	serviceInstanceInit.Do(func() {
		serviceInstance = New(store.Singleton(), standards.RegistrySingleton(), datastore.Singleton())
	})
	return serviceInstance
}
