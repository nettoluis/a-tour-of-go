package main

import "fmt"

func main() {
    pows := make([]int, 10)

    for i := range pows {
        pows[i] = 1 << uint(i) //É a mesma coisa que elevar 2 a i, por conta da estrutura binária e porque o valor padrão é 0
    }

    for _, value := range pows {
        fmt.Printf("%d\n", value)
    }
}
