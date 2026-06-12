package main
import "fmt"

func combination(n, k  int) int {
    if k == 0 || k == n {
        return 1
    }

    return combination(n-1, k-1) + combination(n-1, k)
}

func main() {
    var k, n int

    fmt.Scan(&k, &n)

    fmt.Println(combination(k, n))

}