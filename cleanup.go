package core

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// CleanupFunc is a function that will be called to cleanup resources on application exit
type CleanupFunc func() error

// mutex protects the following variables
var mutex *sync.Mutex
var registered []CleanupFunc

func init() {
	mutex = &sync.Mutex{}
	registered = nil
}

// Cleanup all core services library resources that have been registered by the service.
func Cleanup() error {
	mutex.Lock()
	defer mutex.Unlock()
	if registered == nil {
		return nil
	}
	// Run all registered cleanup funcs
	errs := []string{}
	for _, f := range registered {
		if err := f(); err != nil {
			errs = append(errs, err.Error())
		}
	}
	// Clear registered cleanup funcs
	registered = nil
	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("error(s) occured during core system library cleanup: %v",
			strings.Join(errs, "\n")))
	}
	return nil
}

// RegisterCleanup registers a function to be called at application cleanup
func RegisterCleanup(f CleanupFunc) {
	mutex.Lock()
	defer mutex.Unlock()
	if registered == nil {
		registered = []CleanupFunc{}
	}
	registered = append(registered, f)
}
