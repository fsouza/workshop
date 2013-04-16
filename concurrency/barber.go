// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"time"
)

type Barber int

type Customer int

func (b Barber) CutHair(customers chan Customer) {
	select {
	case c := <-customers:
		fmt.Printf("Cutting hair of customer %d...\n", c)
		time.Sleep(2e9)
	default:
		fmt.Println("Sleeping...")
		time.Sleep(1e9)
	}
}

func main() {
	block := make(chan bool)
	chairs := flag.Int("chairs", 5, "Number of waiting chairs")
	nCustomers := flag.Int("customers", 1200, "Number of customers")
	flag.Parse()
	customers := make(chan Customer, *chairs)
	var barber Barber
	go func(b Barber, customers chan Customer) {
		for {
			b.CutHair(customers)
		}
	}(barber, customers)
	for i := 0; i < *nCustomers; i++ {
		go func(i int) {
			select {
			case customers <- Customer(i):
				fmt.Printf("Customer %d is going to get a haircut...\n", i)
			default:
				fmt.Printf("Customer %d is giving up...\n", i)
			}
		}(i)
		time.Sleep(1e9)
	}
	<-block
}
