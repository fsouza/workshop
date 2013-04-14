// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"
)

type Person struct {
	Name  string
	Birth time.Time
}

func (p *Person) IsBirthday() bool {
	year, month, day := time.Now().Date()
	birthYear, birthMonth, birthDay := p.Birth.Date()
	return year == birthYear && month == && birthMonth && day == birthDay
}

// TODO: embedding
