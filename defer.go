package main

import "fmt"

func main() {
    defer fmt.Println("Ms. Jackson")

    defer fmt.Print("sorry, ")

    fmt.Print("I'm ")
}
