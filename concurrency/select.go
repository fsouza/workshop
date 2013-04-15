// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name    string
	Country string
}

func GetPeopleByCountry(countryName string) ([]Person, error) {
	return nil, nil
}

func main() {
	result := make(chan []Person, 1)
	errChan := make(chan error, 1)
	go func() {
		people, err := GetPeopleByCountry("Brazil")
		if err != nil {
			errChan <- err
		} else {
			result <- people
		}
	}()
	// do things that don't depend on result nor errChan.
	select {
	case people := <-result:
		fmt.Println(people)
	case err := <-errChan:
		log.Fatal(err)
	}
}
