# Gotainer
Dead simple dependency injection container for Golang. This allows you to easily swap implementations of your interfaces with mocks or alternative implementations. Never manually wire your dependencies into constructor calls again.

## Install
```
go get github.com/djlad/gotainer
```

## Example Usage
### Registering Dependencies
In your main function, call a build function that will create your dependencies. For each interface/type your program needs, call register. If the implementation relies on another dependency, pass the container to the constructor. In the constructor, it will get and store the dependencies it needs. Dependencies must be registered before they're requested. So if dependency A (example: Client) depends on B (example: HTTP) register dependency B first.
```
import (
  "github.com/djlad/gotainer/gontainer"
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
### Requesting Dependencies
In a struct's constructor, request all the dependencies it needs from the container. In this example, I'll have Client depend on a Transport interface.
```
type Client struct {
	transport Transport
}

func (c Client) GetData() string {
	return c.transport.get("fake url 1")
}

func NewClient(con gotainer.Container) Client {
	return Client{
		transport: gotainer.Get[Transport](con),
	}
}
```
In the previous section, we used Register to register an HTTP struct to Transport. So, Get[Transport] will return the HTTP struct.
