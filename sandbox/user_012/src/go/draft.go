package main

import "fmt"

type Node struct {
	info rune
	next *Node
	prev *Node
}

type Editor struct {
	head *Node
	cursor *Node
}

func NewEditor() *Editor {
	sentinela := &Node{}
	sentinela.next = sentinela	
	sentinela.prev = sentinela
	return &Editor{head: sentinela, cursor: sentinela}
}

func (edt *Editor) Insert(ch rune) {
    newNode := &Node{info: ch}

    newNode.prev = edt.cursor
    newNode.next = edt.cursor.next

    edt.cursor.next.prev = newNode
	edt.cursor.next = newNode

    edt.cursor = newNode
}

func (edt *Editor) BackSpace() {
	
	if edt.cursor != edt.head {
        target := edt.cursor

	    target.prev.next = target.next
	    target.next.prev = target.prev

        edt.cursor = target.prev
    }

}


func (edt *Editor) Print(){
    if edt.cursor == edt.head {
        fmt.Print("|")
    }

    for i := edt.head.next; i != edt.head; i = i.next {
        fmt.Print(string(i.info))

        if i == edt.cursor {
            fmt.Print("|")
        }
    }

    fmt.Println()
}

func (edt *Editor) MoveLeft() {

    if edt.cursor != edt.head {
        edt.cursor = edt.cursor.prev
    }
}

func (edt *Editor) MoveRight() {

    if edt.cursor != edt.head {
        edt.cursor = edt.cursor.next
    }
}

func (edt *Editor) Delete() {

}

func main(){

    editor := NewEditor()

	var word string

	fmt.Scan(&word)

    for _, letra := range word {
        editor.Insert(letra)
    }

    editor.Print()
	
}