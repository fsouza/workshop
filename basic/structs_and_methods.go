// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name  string
	Birth time.Time
}

func (p *Person) IsBirthday() bool {
	year, month, day := time.Now().Date()
	birthYear, birthMonth, birthDay := p.Birth.Date()
	return year == birthYear && month == birthMonth && day == birthDay
}

type People []Person

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p People) Len() int {
	return len(p)
}

func main() {
	p := Person{Name: "Francisco", Birth: time.Now().Add(-24 * time.Hour)}
	fmt.Println(p.IsBirthday())
}
