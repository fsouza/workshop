// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package safe

import (
	"runtime"
	"sync"
	"testing"
)

func TestNewCounter(t *testing.T) {
	ct := newCounter(2)
	if ct.v != 2 {
		t.Errorf("NewCounter(2), wrong val. Want %d. Got %d.", 2, ct.v)
	}
}

func TestCounterVal(t *testing.T) {
	ct := newCounter(2)
	ct.v = 5
	if ct.val() != 5 {
		t.Errorf("counter.val. Want %d. Got %d.", 5, ct.val())
	}
}

func TestIncrement(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(16))
	n := 50
	var wg sync.WaitGroup
	var ct counter
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < n; i++ {
				ct.increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	want := int64(n * n)
	if ct.val() != want {
		t.Errorf("counter.increment(). Want final %d. Got %d.", want, ct.val())
	}
}

func TestDecrement(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(16))
	n := 50
	var wg sync.WaitGroup
	ct := newCounter(int64(n * n))
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < n; i++ {
				ct.decrement()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if ct.val() != 0 {
		t.Errorf("counter.decrement(). Want final 0. Got %d.", ct.val())
	}
}
