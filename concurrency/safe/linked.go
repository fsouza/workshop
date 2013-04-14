// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package safe

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	length counter
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(v int) {
	node := Node{Value: v, Next: l.Head}
	s := atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.Head)), unsafe.Pointer(node.Next), unsafe.Pointer(&node))
	for !s {
		node.Next = l.Head
		s = atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.Head)), unsafe.Pointer(node.Next), unsafe.Pointer(&node))
	}
	l.length.increment()
}

func (l *LinkedList) Remove() (int, error) {
	if l.length.val() == 0 {
		return 0, errors.New("Empty list.")
	}
	h := l.Head
	s := atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.Head)), unsafe.Pointer(h), unsafe.Pointer(h.Next))
	for !s {
		h = l.Head
		s = atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&l.Head)), unsafe.Pointer(h), unsafe.Pointer(h.Next))
	}
	l.length.decrement()
	return h.Value, nil
}

func (l *LinkedList) Len() int {
	return int(l.length.val())
}
