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

func (v *Set)String() string {
	return "[" + Join(v.data[0:v.size], ", ") + "]"
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
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
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
			v = NewSet(value)
		case "insert":
			// for _, part := range parts[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
			fmt.Println(v)
		case "erase":
			// value, _ := strconv.Atoi(parts[1])
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
