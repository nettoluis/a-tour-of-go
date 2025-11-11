package main

import "fmt"

type Vertex struct {
    X,Y int
}

var (
    v = Vertex{1, 2}
    u = Vertex{X: 12}
    t = Vertex{}
    p = &Vertex{1, 2}
)

func main() {
    fmt.Println(v, u, t, p)
}
