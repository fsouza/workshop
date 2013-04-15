// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
}

func elevator(name string) chan<- Person {
	people := make(chan Person)
	go func() {
		for p := range people {
			fmt.Printf("Elevator %q transporting %q...\n", name, p.Name)
			time.Sleep(1e9)
		}
	}()
	return people
}

func main() {
	people := []Person{
		{Name: "Bob"}, {Name: "Mary"}, {Name: "Thomas"},
		{Name: "John"}, {Name: "Peter"}, {Name: "Ken"},
		{Name: "Patricia"}, {Name: "Ane"}, {Name: "Alice"},
	}
	elevator1 := elevator("social1")
	elevator2 := elevator("social2")
	elevator3 := elevator("vip")
	for _, person := range people {
		select {
		case elevator1 <- person:
		case elevator2 <- person:
		case elevator3 <- person:
		}
	}
	close(elevator1)
	close(elevator2)
	close(elevator3)
}
