package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("Good Morning, Sir")
    case t.Hour() < 17:
        fmt.Println("Good Afternoon, Sir")
    default:
        fmt.Println("Good Evening, Sir")
    }
}
