package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Root struct {
	info int
	next *Root
	prev *Root
}

type Dlist struct {
	head *Root
}

func NewLList() *Dlist {
	dlist := Dlist{}
	dlist.head = &Root{}
	dlist.head.next = dlist.head
	dlist.head.prev = dlist.head

	return &dlist
}

func insert(ref *Root, value int) {
	anterior := ref.prev
	new := &Root {
		info: value,
		next: ref,
		prev: anterior,
	}

	anterior.next = new
	ref.prev = new
}


func (list *Dlist) PushBack(value int) {
	insert(list.head, value)
}

func (list *Dlist) PushFront(value int) {
	insert(list.head.next, value)
}
func (list *Dlist) Front() *Root {
	return list.head.next
}

func (list *Dlist) Back() *Root {
	return list.head.prev
}

func (list *Dlist) End() *Root {
	return list.head
}

func (list *Dlist) Size() int{
	var count int

	for it := list.Front(); it != list.head; it = it.next {
		count++
	}

	return count
}

func (list *Dlist) Clear() {
	list.head.next = list.head
	list.head.prev = list.head
}

func (list *Dlist) PopFront() {
	if list.head.next == list.head {
		return
	}

	removido := list.head.next
	proximo := removido.next

	list.head.next = proximo
	proximo.prev = list.head

	removido.next = nil 
	removido.prev = nil 
}

func (list *Dlist) PopBack() {
	removido := list.head.prev
	ultimo := removido.prev

	ultimo.next = list.head
	list.head.prev = ultimo

}
func (list *Dlist) String() string {
	saida := "["
	for it := list.Front(); it != list.head; it = it.next {
		saida += fmt.Sprintf("%d", it.info)

		if it.next != list.head {
			saida += ", "
		}
	}

	return saida + "]"
}


func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
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
			fmt.Println(ll.Size())
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
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
