// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func apply(values []int, fn func(int)) {
	for _, v := range values {
		fn(v)
	}
}

// count people in the database
func countPeople() (int, error) {
	n := 10
	return n, nil
}

func getInfo() (result int, err error) {
	result = 10
	return
}

func main() {
	values := []int{1, 2, 3}
	apply(values, func(i int) { fmt.Println(i) })
}
