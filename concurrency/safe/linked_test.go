// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package safe

import (
	"sync"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	linked := NewLinkedList()
	if linked.Len() != 0 {
		t.Errorf("NewLinkedList(): Want 0. Got %d.", linked.Len())
	}
}

func TestLinkedListInsert(t *testing.T) {
	values := []int{10, 15, 18}
	linked := NewLinkedList()
	for _, v := range values {
		linked.Insert(v)
	}
	n := linked.Head
	if n == nil {
		t.Fatal("Got unexpected <nil> on LinkedList.Head")
	}
	for i := len(values) - 1; n != nil; i-- {
		if n.Value != values[i] {
			t.Errorf("LinkedList.Insert(). Want %d. Got %d.", values[i], n.Value)
		}
		n = n.Next
	}
	if linked.Len() != 3 {
		t.Errorf("LinkedList.Len(): Want 3. Got %d.", linked.Len())
	}
}

func TestLinkedListRemove(t *testing.T) {
	values := []int{10, 15, 18}
	linked := NewLinkedList()
	for _, v := range values {
		linked.Insert(v)
	}
	for i := len(values) - 1; i >= 0; i-- {
		v, err := linked.Remove()
		if err != nil {
			t.Fatal(err)
		}
		if v != values[i] {
			t.Errorf("LinkedList.Remove(). Want %d. Got %d.", values[i], v)
		}
		if linked.Len() != i {
			t.Errorf("LinkedList.Len(): Want %d. Got %d.", i, linked.Len())
		}
	}
}

func TestLinkedListRemoveEmpty(t *testing.T) {
	l := NewLinkedList()
	_, err := l.Remove()
	if err == nil {
		t.Error("Want non-nil error, got <nil>.")
	}
}

// Cannot fail with -race.
func TestIsSafe(t *testing.T) {
	linked := NewLinkedList()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			linked.Insert(i)
			wg.Done()
		}(i)
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			linked.Remove()
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkInsert(b *testing.B) {
	linked := NewLinkedList()
	for i := 0; i < b.N; i++ {
		linked.Insert(i * 100)
	}
}

func BenchmarkRemove(b *testing.B) {
	b.StopTimer()
	linked := NewLinkedList()
	for i := 0; i < 1e6; i++ {
		linked.Insert(i * 100)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		linked.Remove()
	}
}
