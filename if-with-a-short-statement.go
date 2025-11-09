package main

import (
    "fmt"
    "math"
)

func pow(m, n, lim float64) float64 {
    if v := math.Pow(m, n); v < lim {
        return v        
    }
    return lim
}

func main() {
    fmt.Println(
        pow(2, 5, 12),
        pow(1, 523, 2),
    )
}
