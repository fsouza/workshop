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

func main() {
	values := []int{1, 2, 3}
	apply(values, func(i int) { fmt.Println(i) })
}
