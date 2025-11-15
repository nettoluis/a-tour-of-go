package main

import "fmt"

func main() {
    a := []int{1, 2, 3, 4, 5, 6}
    fmt.Println(a)

    s := a[:0]
    fmt.Println(s)
    fmt.Printf("Capacidade: %d\nTamanho: %d\n", cap(s), len(s))

    s = s[:4]
    fmt.Println(s)
    fmt.Printf("Capacidade: %d\nTamanho: %d\n", cap(s), len(s))

    s = s[2:]
    fmt.Println(s)
    fmt.Printf("Capacidade: %d\nTamanho: %d\n", cap(s), len(s))

    s = s[:]
    fmt.Println(s)
    fmt.Printf("Capacidade: %d\nTamanho: %d\n", cap(s), len(s))
}
