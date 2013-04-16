// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"sync"
)

var hits = struct {
	m map[string]int
	sync.Mutex
}{m: make(map[string]int)}

func hit(name string) {
	hits.Lock()
	hits.m[name] += 1
	hits.Unlock()
}

func main() {

}
