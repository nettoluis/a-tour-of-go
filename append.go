package main

import "fmt"

func main() {
    var s []int
    printSlice(s)

    s = append(s, 0)
    printSlice(s)

    s = append(s, 12)
    printSlice(s)

    for i := 0; i < 15; i++ {
        s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10) //O tamanho do novo array realocado é sempre o dobro do anterior
        printSlice(s) 
    }
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
