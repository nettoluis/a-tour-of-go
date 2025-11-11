package main

import "fmt"

func main() {
    i, j := 12, 11

    p := &i
    fmt.Println(*p)
    *p = 10
    fmt.Println(*p)

    p = &j
    fmt.Println(*p)
    *p = 15
    fmt.Println(*p)
}
