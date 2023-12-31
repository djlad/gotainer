// Gotainer is a simple dependency injection container for Go.
//
// It can be used to register dependencies to interfaces and then retrieve them later
// It is basically a map between interfaces and implementations.
// Use it to easily swap ouy implementations or mocks.
package main

import (
	"fmt"

	"github.com/djlad/gotainer/gotainer"
)

type Transport interface {
	get(url string) string
}

type HTTP struct {
}

func (h HTTP) get(url string) string {
	return "HTTP get from " + url
}

type SSH struct {
}

func (h SSH) get(url string) string {
	return "SSH get from " + url
}

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

type Client2 struct {
	transport Transport
}

func NewClient2(con gotainer.Container) Client2 {
	return Client2{
		transport: gotainer.Get[Transport](con),
	}
}

func (c Client2) GetData() string {
	return c.transport.get("fake url 2")
	// return "fake url 2"
}

type Driver struct {
	client  Client
	client2 Client2
}

func (d Driver) StartProgram() {
	fmt.Println("Start program")
	fmt.Println(d.client.GetData())
	fmt.Println(d.client2.GetData())
	fmt.Println(d.client.GetData())
	fmt.Println(d.client2.GetData())
}

func NewDriver(con gotainer.Container) Driver {
	return Driver{
		client:  gotainer.Get[Client](con),
		client2: gotainer.Get[Client2](con),
	}
}

func build() gotainer.Container {
	container := gotainer.NewContainer()
	gotainer.Register[Transport](container, HTTP{})
	// gotainer.Register[Transport](container, SSH{})
	gotainer.Register(container, NewClient(container))
	gotainer.Register(container, NewClient2(container))
	gotainer.Register(container, NewDriver(container))
	return container
}

func main() {
	con := build()
	driver := gotainer.Get[Driver](con)
	driver.StartProgram()
}
