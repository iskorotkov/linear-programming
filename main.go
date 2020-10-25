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

	geometricSolution(a)
}

func geometricSolution(a vectors.Matrix) {
	findMainVariables(a)
	makeMatrixDiagonal(a)
	x1, x2, x3 := findXFromDiagonalMatrix(a)
	c := expressMainVariables(x1, x2, x3)
	checkThatXIsCorrect(a, x1, x2, x3)
	x4, x5 := findFreeVariables(x1, x2, x3)

	c.ReplaceVariable(0, x1)
	c.ReplaceVariable(1, x2)
	c.ReplaceVariable(2, x3)
	c.ReplaceVariable(3, x4)
	c.ReplaceVariable(4, x5)
	fmt.Printf("c_min = %v\n", c[len(c)-1])
}

func findFreeVariables(x1, x2, x3 vectors.Vector) (vectors.Vector, vectors.Vector) {
	x4 := make(vectors.Vector, 6)
	x5 := make(vectors.Vector, 6)
	copy(x5, x2)
	x5.ReplaceVariable(3, x4)
	x5 = x5.GetVariable(4)
	fmt.Printf("x4 = %v\nx5 = %v\n", x4, x5)
	return x4, x5
}

func checkThatXIsCorrect(a vectors.Matrix, x1 vectors.Vector, x2 vectors.Vector, x3 vectors.Vector) {
	for i := range a {
		a[i].ReplaceVariable(0, x1)
		a[i].ReplaceVariable(1, x2)
		a[i].ReplaceVariable(2, x3)
	}
	fmt.Printf("should be zero:\n%v\n", a)
}

func expressMainVariables(x1 vectors.Vector, x2 vectors.Vector, x3 vectors.Vector) vectors.Vector {
	c := vectors.Vector{8, 10, -1, -3, -2, 0}
	c.ReplaceVariable(0, x1)
	c.ReplaceVariable(1, x2)
	c.ReplaceVariable(2, x3)
	fmt.Printf("c = %v\n", c)
	return c
}

func findXFromDiagonalMatrix(a vectors.Matrix) (vectors.Vector, vectors.Vector, vectors.Vector) {
	x3 := a.GetVariable(2)

	x2 := a.GetVariable(1)
	x2.ReplaceVariable(2, x3)

	x1 := a.GetVariable(0)
	x1.ReplaceVariable(1, x2)
	x1.ReplaceVariable(2, x3)

	fmt.Printf("x1 = %v\n", x1)
	fmt.Printf("x2 = %v\n", x2)
	fmt.Printf("x3 = %v\n", x3)
	return x1, x2, x3
}

func makeMatrixDiagonal(a vectors.Matrix) {
	a.SubtractRow(0)
	a[1].MultiplyBy(3)
	a.SubtractRow(1)
	fmt.Printf("a:\n%v\n", a)
}

func findMainVariables(a vectors.Matrix) {
	a1 := a.Slice(vectors.Range{0, 3}, vectors.Range{0, 3})
	fmt.Printf("a1:\n%v\n", a1)
	fmt.Printf("|a1| = %v\n", a1.Determinant())
}
