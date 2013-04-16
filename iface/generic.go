// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func main() {
	ch := make(chan int)
	var values []interface{}
	values = append(values, 10, "chico", 80.0, ch)
	values = append(values, values)
	values = append(values, values...)
	fmt.Println(values)
}
