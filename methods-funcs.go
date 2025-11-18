package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(x.X * v.Y + v.X * v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
