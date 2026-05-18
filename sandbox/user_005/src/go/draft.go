    package main
    import "fmt"

    type Multiset struct {
        data []int
        size int
        capacity int
    }

    func NewMultiSet(capacity int) *Multiset{
        return &Multiset {
            data: make([]int, capacity),
            size: 0,
            capacity: capacity,
        }
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
            fmt.Printf("O VALOR %d EXISTE\n", value)
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

    func (ms *Multiset) Insert(value int) {
        i := ms.size - 1

        for i >=0 && ms.data[i] > value {
            i--
        }

        index_insertion := i + 1
        
        ms.insert(value, index_insertion)
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
            return fmt.Errorf("value not found")
        }

        return ms.erase(index)
    }

    func (ms *Multiset) erase(index int) error {

        if index > ms.capacity || index < 0{
            return fmt.Errorf("value note found")
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

    func main() {
        ms := NewMultiSet(0)

        ms.expand()
        ms.Insert(1)
        ms.Insert(2)
        ms.Insert(3)
        ms.Insert(1)

        ms.Contains(1)

        fmt.Println(ms.data)

        ms.Erase(1)

        fmt.Println(ms.data[:ms.size])

        ms.Erase(4)
    }
