// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
}

type Elevator struct {
	name   string
	people <-chan Person
}

func NewElevator(name string, people <-chan Person) *Elevator {
	return &Elevator{name: name, people: people}
}

func (e *Elevator) Start() {
	go func() {
		for p := range e.people {
			fmt.Printf("%q transporting %s.\n", e.name, p.Name)
			time.Sleep(time.Duration((rand.Int()%5 + 1) * 1e9))
		}
	}()
}

func main() {
	nPeople := flag.Int("people", 100, "Number of people")
	nElevators := flag.Int("elevators", 4, "Number of elevators")
	flag.Parse()
	people := make(chan Person)
	for i := 0; i < *nElevators; i++ {
		name := fmt.Sprintf("elevator %d", i+1)
		e := NewElevator(name, people)
		e.Start()
	}
	for i := 0; i < *nPeople; i++ {
		name := fmt.Sprintf("person %d", i+1)
		people <- Person{Name: name}
	}
	close(people)
}
