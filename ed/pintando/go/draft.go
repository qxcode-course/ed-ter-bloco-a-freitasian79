package main
import (
    "fmt"
    "math"
)
func main() {
    var num1, num2, num3 float64
    
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    fmt.Scan(&num3)

    var p = (num1 + num2 + num3) / 2

    var area = math.Sqrt(p * (p - num1) * (p - num2) * (p - num3))

    fmt.Printf("%.2f\n",area)

}
