package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    saida := make([][]uint8, dy)

    for i := range saida {
        linha := make([]uint8, dx)
        for j := range linha {
            linha[j] = uint8((i * j) * (i * j))
        }

        saida[i] = linha
    }

    return saida
}

func main() {
    pic.Show(Pic)
}

