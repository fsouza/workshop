// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"
)

func greetings(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	var once sync.Once
	for i := 0; i < 10000; i++ {
		once.Do(func() {
			greetings("Gopher")
		})
	}
}
