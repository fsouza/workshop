// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Test struct {
	Theme string
}

func student(test Test, c *sync.Cond) {
	fmt.Println("getting ready...")
	c.L.Lock()
	c.Wait()
	fmt.Println("doing test")
	time.Sleep(1e9)
	c.L.Unlock()
}

func main() {
	var mut sync.Mutex
	cond := sync.NewCond(&mut)
	for i := 0; i < 10; i++ {
		go student(Test{Theme: "math"}, cond)
	}
	time.Sleep(5e9)
	cond.Broadcast()
	time.Sleep(20e9)
}
