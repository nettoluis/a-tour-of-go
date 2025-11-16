package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["nome"] = 1
    fmt.Println("O valor é ", m["nome"])

    m["nome"] = 2
    fmt.Println("O valor é ", m["nome"])

    delete(m, "nome")
    fmt.Println("O valor é ", m["nome"])

    v, ok := m["nome"]
    fmt.Println("O valor é ", v, "Presente?", ok)
}


