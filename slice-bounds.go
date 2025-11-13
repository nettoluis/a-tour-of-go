package main

import "fmt"

func main() {
    s := [12]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

    r := s[:3]
    fmt.Println(r)

    u := s[4:9]
    fmt.Println(u)

    t := s[:]
    fmt.Println(t)
}
