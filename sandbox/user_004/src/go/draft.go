package main
import "fmt"

type Node struct {
    info int
    next *Node
    prev *Node
}

type DList struct {
    head *Node
}

func NewDList() *DList {
    dlist := DList{}
    dlist.head = &Node{}
    dlist.head.next = dlist.head
    dlist.head.prev = dlist.head

    return &dlist
}

func insert(ref *Node, value int) {
    previous := ref.prev

    new_node := &Node {
        info: value,
        next: ref,
        prev: previous,
    }

    previous.next = new_node
    ref.prev = new_node
}

func (list *DList) PushBack(value int) {
    insert(list.head, value)
}

func (list *DList) PushFront(value int) {
    insert(list.head.next, value)
}

func show(list *DList) {
    it := list.head.next 

    fmt.Print("[")
    for it = list.head.next; it != list.head; it = it.next {
        fmt.Printf(" %d", it.info)
    }
    fmt.Print(" ]\n")
}
//  AMBAS ESTÃO FUNCIONANDO ****
func (list *DList) Clear() {
    list.head.next = list.head
    list.head.prev = list.head
}

func Clear(list *DList) {
    list.head.next = list.head
    list.head.prev = list.head
}
// *****************************

func (list *DList) PopBack(){
    removed := list.head.prev
    new_last := removed.prev

    new_last.next = list.head
    list.head.prev = new_last

}

func main(){
    list := NewDList()

    for i := range 10 {
        list.PushFront(i)
    }
    
    show(list)
    list.PopBack()
    show(list)
}


