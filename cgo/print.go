// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"

import "unsafe"

func main() {
	str := C.CString("Hello, world!\n")
	defer C.free(unsafe.Pointer(str))
	C.printf(str)
}
