# Gotainer
Dead simple dependency injection container for Golang. Easily swap implementations of your interfaces with mocks or alternative implementations.

## Install
```
go get github.com/djlad/gotainer
```

## Example Usage
```
import (
  "github.com/djlad/gotainer/container"
  ...
)

func build() gotainer.Container {
  // Registers a dependency struct HTTP to interface Transport
  gotainer.Register[Transport](container, HTTP{})
  // Registers a dependency Client struct as a singleton to the type Client
  // Client constructor takes container so it can request Transport interface dependency
  // (in this case the implementation will be HTTP)
  gotainer.Register[Client](container, NewClient(container))
  // register a driver that will Get its dependencies.
  // explicitly passing the type to register to is unnecessary since
  // the type will be inferred by default. But, I'll add the generic paramater for clarity.
  // If I was registering to an interface, then I'd have to specify the interface as a parameter
  gotainer.Register[Driver](container, NewDriver(container))
  return container
}

func main(){
  // build your container
  container := build()
  // Get any dependency. In this case we'll get driver to start the program.
  driver := gotainer.Get[Driver](container)
  driver.start()
}
```
