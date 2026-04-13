package main
import "fmt"

func process(value int) {

    if (value == 0) {
        return
    }

    div := value / 2
    rest := value % 2
    
    process(div)

    fmt.Printf("%d %d\n", div, rest)
        
}

func main() {
    var value int
    fmt.Scan(&value)
    process(value)

}
