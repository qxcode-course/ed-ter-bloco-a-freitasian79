package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Set struct {
	data []int
	size int
	capacity int
}
func NewSet(capacity int) *Set {
	return &Set {
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (st *Set)String() string {
	return "[" + Join(st.data[0:st.size], ", ") + "]"
}


func (st *Set) Insert(value int)  {
	index := 0

	for index < st.size {
		if st.data[index] == value {
			return
		} 

		if st.data[index] > value {
			break
		}

		index++
	}

	if st.size == st.capacity {
		newCap := st.capacity * 2
		if newCap == 0 {
			newCap = 1
		}	
		
		st.reserve(newCap)
	}

	st.insert(value, index)

}

func (st *Set) reserve(newCapacity int) {
	if newCapacity <= st.capacity {
		return
	}

	newSet := make([]int, newCapacity)

	copy(newSet, st.data)
	st.data = newSet
	st.capacity = newCapacity
}

func (st *Set) insert(value int, index int) error {
	if st.capacity < st.size {
		return fmt.Errorf("index out of range")
	}

	for i := st.size; i > index; i-- {
		st.data[i] = st.data[i - 1]
	}

	st.data[index] = value
	st.size++

	return nil
}

func (st *Set) Contains(value int) bool {
	indice := st.binarySearch(value)

	if indice != -1 {
		return true
	}

	return false
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

func (st *Set) binarySearch(value int) int {
	start := 0
	end := st.size - 1

	for start <= end {
		middle := (start + end) / 2

		if value == st.data[middle] {
			return middle
		}

		if value > st.data[middle] {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}

func (st *Set) Erase(value int) bool {
	if st.binarySearch(value) != -1 {
		valor_removido := st.binarySearch(value)
		st.erase(valor_removido)
		return true
	}

	return false
}

func (st *Set) erase(index int) error {
	if index > st.size || index < st.size {
		fmt.Errorf("value not found")
	}

	for i := index; i < st.size - 1; i++ {
		st.data[i] = st.data[i + 1]
	}

	st.size--

	return nil
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	st := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			st = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				st.Insert(value)
			}
		case "show":
			fmt.Println(st.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !st.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(st.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
