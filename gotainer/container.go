package gotainer

import (
	"fmt"
	"reflect"
)

// Container is a dependency injection container. Pass it to the constructors of your structs and they
// can use the Get function to request the dependencies they need. Then register them to the container for an interface
// and the next dependencies can request the structs you registered. This will make it easy to swap out
// the implementations of your interfaces with mocks or alternative implementations.
type Container struct {
	dependencies map[reflect.Type]any
}

// RegisterFactory registers a factory function to an interface or type. Whenever this dependency
// is requested by Get, the factory function will be called and the result will be returned.
// If called twice for the same interface or type, the second call will overwrite the first.
func RegisterFactory[DependencyType any](container Container, factory func() DependencyType) {
	t := reflect.TypeOf((*DependencyType)(nil)).Elem()
	container.dependencies[t] = factory
}

// Register registers a singleton to an interface or type. Whenever this dependency
// is requested by Get, the same instance will be returned.
// If called twice for the same interface or type, the second call will overwrite the first.
func Register[DependencyType any](container Container, singleton DependencyType) {
	t := reflect.TypeOf((*DependencyType)(nil)).Elem()
	container.dependencies[t] = func() DependencyType { return singleton }
}

// Get retrieves the dependency stored by Register or RegisterFactory
// that implments the type or interface given as a generic parameter.
// Panics if no dependency was registered.
func Get[DependencyType any](container Container) DependencyType {
	dummy := (*DependencyType)(nil)
	t := reflect.TypeOf(dummy).Elem()
	dep, ok := container.dependencies[t]
	if !ok {
		panic(fmt.Sprintf("Dependency %s not found", t))
	}
	return dep.(func() DependencyType)()
}

func NewContainer() Container {
	return Container{
		dependencies: make(map[reflect.Type]any),
	}
}
