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

func elevator() chan<- Person {
	people := make(chan Person, 2)
	go func() {
		for p := range people {
			fmt.Printf("Transporting %q...\n", p.Name)
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
	ch := elevator()
	for _, person := range people {
		ch <- person
	}
	close(ch)
}
