package main

import "fmt"

func main() {
    contador := 5
    for {
        if contador == 5 {
            break
        }
        contador += contador
    }

    fmt.Println(contador)
}
