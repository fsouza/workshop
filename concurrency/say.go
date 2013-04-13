// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

func say(what string) {
	fmt.Println(what)
}

func main() {
	go say("hello world")
}
