package main

import "fmt"

func main() {
    names := [4]string{
        "John",
        "Paul",
        "Ringo",
        "George",
    }

    a := names[0:2]
    b := names[2:]
    fmt.Println(a, b)

    b[0] = "Eu né"
    fmt.Println(a, b)
    fmt.Println(names)
}
