package main

import "fmt"

func main() {
    a := make([]int, 5) //define o len como 5
    b := make([]int, 0, 5) //define o len como 0 e a cap como 5

    fmt.Printf("A: %s\n", sliceStatus(a))
    fmt.Printf("B: %s\n", sliceStatus(b))

    fmt.Println(a,b)
}

func sliceStatus (a []int) string {
    saida := fmt.Sprintf("Capacidade: %d\nTamanho: %d", cap(a), len(a))

    return saida
}
