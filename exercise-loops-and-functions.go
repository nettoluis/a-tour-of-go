package main

import (
    "fmt"
    "math"
)

func abs(x float64) float64 {
    return math.Pow(math.Pow(x, 2), 0.5)
}

func sqrt(x float64) float64 {
    margin := 1e-10
    if x <= 0 {
        return 0.0
    }

    z := float64(x)
    for {
        new_z := 0.5 * (z + x / z)
        if abs(new_z - z) < margin {
            return new_z
        }
        z = new_z
    }
}

func main() {
    fmt.Println(sqrt(2))
    fmt.Println(math.Sqrt(2))
    fmt.Println(sqrt(4))
    fmt.Println(math.Sqrt(4))
    fmt.Println(sqrt(25))
    fmt.Println(math.Sqrt(25))
}
