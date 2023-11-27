// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Hello is a simple hello, world demonstration web server.
//
// It serves version information on /version and answers
// any other request like /name by saying "Hello, name!".
package main

import (
	"fmt"

	"github.com/djlad/gotainer/container"
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

func NewClient(con container.Container) Client {
	return Client{
		transport: container.Get[Transport](con),
	}
}

type Client2 struct {
	transport Transport
}

func NewClient2(con container.Container) Client2 {
	return Client2{
		transport: container.Get[Transport](con),
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

func NewDriver(con container.Container) Driver {
	return Driver{
		client:  container.Get[Client](con),
		client2: container.Get[Client2](con),
	}
}

func main() {
	con := container.NewContainer()
	// container.Register[Transport](con, func() Transport { return HTTP{} })

	container.Register[Transport](con, SSH{})
	container.Register[Transport](con, HTTP{})
	container.Register(con, NewClient(con))
	container.Register(con, NewClient2(con))
	container.Register(con, NewDriver(con))
	driver := container.Get[Driver](con)
	driver.StartProgram()
}
