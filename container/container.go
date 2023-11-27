package container

import (
	"reflect"
)

type Container struct {
	dependencies map[reflect.Type]any
}

func RegisterFactory[DependencyType any](container Container, factory func() DependencyType) {
	t := reflect.TypeOf((*DependencyType)(nil)).Elem()
	container.dependencies[t] = factory
}

func Register[DependencyType any](container Container, singleton DependencyType) {
	t := reflect.TypeOf((*DependencyType)(nil)).Elem()
	container.dependencies[t] = func() DependencyType { return singleton }
}

func Get[DependencyType any](container Container) DependencyType {
	dummy := (*DependencyType)(nil)
	t := reflect.TypeOf(dummy).Elem()
	dep := container.dependencies[t].(func() DependencyType)
	return dep()
}

func NewContainer() Container {
	return Container{
		dependencies: make(map[reflect.Type]any),
	}
}
