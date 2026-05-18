package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Multiset struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *Multiset {
	return &Multiset {
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *Multiset) Insert(value int) {
	i := ms.size - 1

	for i >= 0 && ms.data[i] > value {
		i--
	}

	index_insertion := i + 1

	ms.insert(value, index_insertion)
}

func (ms *Multiset) insert(value int, index int) error {
	if index > ms.capacity || index < 0 {
		fmt.Errorf("index out of range")
	}

	if ms.size >= ms.capacity {
		ms.expand()
	}

	for i := ms.size - 1; i >= index; i-- {
		ms.data[i + 1] = ms.data[i]
	}

	ms.data[index] = value
	ms.size++

	return nil
}

func (ms *Multiset) Contains(value int) bool {
	if ms.binarySearch(value) != -1 {
		return true
	}

	return false
}

func (ms *Multiset) binarySearch(value int) int {
	start := 0
	end := ms.size - 1

	for start <= end {
		middle := (start + end) / 2

		if ms.data[middle] == value {
			return middle
		} else if ms.data[middle] < value {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}

func (ms *Multiset)String() string {
	return "[" + Join(ms.data[0:ms.size], ", ") + "]"
}

func (ms *Multiset) expand() {
	if ms.capacity == 0 {
		ms.capacity = 1
	} else {
		ms.capacity *= 2
	}

	new_data := make([]int, ms.capacity)
	copy(new_data, ms.data)
	ms.data = new_data
}

func (ms *Multiset) Erase(value int) error {
	
	index := ms.binarySearch(value)

	if index == -1 {
		return fmt.Errorf("value note found")
	}

	return ms.erase(index)
}

func (ms *Multiset) erase(index int) error {
	
	if index > ms.capacity || index < 0 {
		return fmt.Errorf("value not found")
	}

	for i := index; i < ms.size - 1; i++ {
		ms.data[i] = ms.data[i + 1]
	}

	ms.size--
	ms.data[ms.size] = 0
	
	return nil
}

func (ms *Multiset) Count(value int) int {
        
    counter := 0

    for i := 0; i < ms.size; i++ {
        if ms.data[i] == value {
			counter++
        } else if ms.data[i] > value {
            break
        }
    }

        return counter
}

func (ms *Multiset) Unique() int {
	if ms.size == 0 {
		return 0
	}
	counter := 1

	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i - 1] {
			counter++
		}
	}

	return counter 
}

func (ms *Multiset) Clear() {
	ms.data = make([]int, 0)
	ms.capacity = 0
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value) 
			if err != nil {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
		value, _ := strconv.Atoi(args[1])
		fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
