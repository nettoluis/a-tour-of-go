package main

import "fmt"

func main() {
    pows := []int{1, 2, 4, 8, 16, 32}

    for index, value := range pows {
        fmt.Printf("2 elevado a %d é igual a %d\n", index, value)
    }
}
