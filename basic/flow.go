// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

func countPeople() (int, error) {
	return 10, nil
}

func getValue() int {
	return 5
}

func main() {
	n, err := countPeople()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
	if x, err := countPeople(); err == nil {
		fmt.Println(x)
	} else {
		log.Fatal(err)
	}
	fmt.Println(x) // BOOM

	v := getValue()
	switch v {
	case 10:
		fmt.Println("something")
		// no break
	case 5:
		fmt.Println("otherthing")
	default:
		fmt.Println("what")
	}

	switch {
	case v < 4 || v > 8:
		fmt.Println("not in range")
	case v > 6:
		fmt.Println("high")
	case v < 6:
		fmt.Println("low")
	default:
		fmt.Println("six")
	}
}
