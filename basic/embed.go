// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Widget struct {
	X, Y int
}

func (w *Widget) Dimensions() (int, int) {
	return w.X, w.Y
}

type Label struct {
	Widget
	Text string
}

func main() {
	l := Label{
		Widget: Widget{X: 10, Y: 10},
		Text:   "Name",
	}
	l.Dimensions()
}
