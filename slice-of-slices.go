package main

import (
    "fmt"
    "strings"
)

func main() {
    board := [][]string {
        []string{"-", "-", "-"},
        []string{"-", "-", "-"},
        []string{"-", "-", "-"},
    }

    board[0][0] = "X"
    board[0][1] = "0"
    board[1][1] = "X"
    board[0][2] = "0"
    board[2][2] = "X"

    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }
}
