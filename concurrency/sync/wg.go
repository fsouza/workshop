// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	names := []string{"Bob", "Alice", "Ana", "Mary"}
	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			fmt.Printf("Hello, %s!\n", name)
			wg.Done()
		}(name)
	}
	wg.Wait()
}
