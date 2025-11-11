package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = ", World!"
    p1 := &a[0]
    p2 := &a[1]
    fmt.Println(*p1, *p2)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)
}
