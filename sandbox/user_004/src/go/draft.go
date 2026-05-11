package main
import "fmt"

type Root struct {
    info int
    next *Root
    prev *Root
}

type Dlist struct {
    head *Root
}

func NewList() *Dlist {
    dlist := Dlist{}
    dlist.head = &Root{}
    dlist.head.next = dlist.head
    dlist.head.prev = dlist.head

    return &dlist
}

func insert(ref *Root, value int) {
    anterior := ref.prev

    C := &Root {
        info: value,
        next: ref,
        prev: anterior,
    }

    anterior.next = C
    ref.prev = C

}

func PushBack(list *Dlist, value int) {
    insert(list.head, value)
}

func PushFront(list *Dlist, value int) {
    insert(list.head.next, value)
}

func print(list *Dlist) {
    it := list.head.next
    fmt.Print("[")
    for {
        if it == list.head {
            break

        }
        fmt.Printf(" %d", it.info)
        it = it.next
    }

    fmt.Print(" ]\n")
}
func main() {
    lista := NewList()

    for i := range 10 {
        PushBack(lista, i)
        print(lista)
        fmt.Print("* =============== *\n")
    }

    print(lista)
}
