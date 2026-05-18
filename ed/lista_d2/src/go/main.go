package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int 
	next *Node
	prev *Node
	root *Node
}

type Dlist struct {
	root *Node
	size int
}

func NewLList() *Dlist {
	dlist := Dlist{}
	dlist.root = &Node{}
	dlist.root.next = dlist.root
	dlist.root.prev = dlist.root

	return &dlist
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}

	return n.prev
}

func (list *Dlist) insert(ref *Node, value int) {
	anterior := ref.prev

	new := &Node {
		Value: value,
		next: ref,
		prev: anterior,
		root: list.root,
	}

	anterior.next = new
	ref.prev = new
}

func (list *Dlist) Front() *Node {
	if list.root.next == list.root {
		return nil
	}

	return list.root.next
}

func (list *Dlist) Back() *Node {
	if list.root.prev == list.root {
		return nil
	}

	return list.root.prev
}

func (list *Dlist) End() *Node {
	return list.root
}

func (list *Dlist) PushBack(value int) {
	list.insert(list.root, value)
}

func (list *Dlist) PushFront(value int) {
	list.insert(list.root.next, value)
}

func (list *Dlist) Clear() {
	list.root.next = list.root
	list.root.prev = list.root
}
func (list *Dlist) String() string {
	out := "["

	for it := list.root.next; it != list.root; it = it.next {
		out += fmt.Sprintf("%d", it.Value)

		if it.next != list.root {
			out += ", "
		}
	}

	return out + "]"
}

func (list *Dlist) Search(value int) *Node {
	for it := list.root.next; it != list.root; it = it.next {
		if it.Value == value {
			return it
		}
	}

	return nil
}

func (list *Dlist) Insert(ref *Node, value int) {
	if ref == nil {
		return 
	}

	list.insert(ref, value)
}

func (list *Dlist) Remove(remv *Node) {
	anterior := remv.prev
	posterior := remv.next

	anterior.next = posterior
	posterior.prev = anterior
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
			 	ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
			 	num, _ := strconv.Atoi(v)
			 	ll.PushFront(num)
			 }
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
			 	fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			 oldvalue, _ := strconv.Atoi(args[1])
			 newvalue, _ := strconv.Atoi(args[2])
			 node := ll.Search(oldvalue)
			 if node != nil {
			 	node.Value = newvalue
			 } else {
			 	fmt.Println("fail: not found")
			 }
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
