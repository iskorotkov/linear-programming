package main

import (
	"fmt"

	"github.com/iskorotkov/linear_programming/vectors"
)

func main() {
	a := vectors.Matrix{
		{6, 5, 7, 10, 7, -71},
		{4, 3, 6, 6, 3, -45},
		{6, 6, 8, 8, -2, -54},
	}

	a1 := a.Slice(vectors.Range{0, 3}, vectors.Range{0, 3})
	fmt.Printf("|a1| = %v\n", a1.Determinant())

	a2 := a.Slice(vectors.Range{0, 3}, vectors.Range{1, 4})
	fmt.Printf("|a2| = %v\n", a2.Determinant())

	a.SubtractRow(0)
	a[1].MultiplyBy(3)
	a.SubtractRow(1)
	fmt.Printf("a:\n%v\n", a)

	x3 := a.GetVariable(2)

	x2 := a.GetVariable(1)
	x2.ReplaceVariable(2, x3)

	x1 := a.GetVariable(0)
	x1.ReplaceVariable(1, x2)
	x1.ReplaceVariable(2, x3)

	fmt.Printf("x1 = %v\n", x1)
	fmt.Printf("x2 = %v\n", x2)
	fmt.Printf("x3 = %v\n", x3)

	c := vectors.Vector{8, 10, -1, -3, -2, 0}
	c.ReplaceVariable(0, x1)
	c.ReplaceVariable(1, x2)
	c.ReplaceVariable(2, x3)
	fmt.Printf("c = %v\n", c)
}
