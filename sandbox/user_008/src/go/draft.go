package main

import (
	"fmt"
	"lista"
)

type Node[T comparable] struct {
	Value T
	next *Node[T]
	prev *Node[T]
	root *Node[T]
}

type LList struct {
	root *Node[T]
	size int
}

func NewLList[T comparable]() *LList[T] {
	root := &Node[T]{}
	root.next = root
	root.prev = root
	root.root = root

	return &LList[T]{root: root, size: 0}
}

func main() {
    
}
