package main

import "fmt"

func main() {
    q := []int{12, 1, 4}
    fmt.Println(q)

    r := []bool{true, true, true}
    fmt.Println(r)

    s := []struct {
        i int
        b bool
    }{
        {1, true},
        {2, true},
        {3, true},
    }

    fmt.Println(s)


}
