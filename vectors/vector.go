package vectors

import (
	"fmt"
)

type Vector []float64

func (v Vector) Length() int {
	return len(v)
}

func (v Vector) GetVariable(index int) Vector {
	coef := v[index]
	res := make([]float64, 0, v.Length())
	for _, value := range v {
		res = append(res, -value/coef)
	}
	res[index] = 0
	return res
}

func (v Vector) MultiplyBy(multiplier float64) {
	for i := range v {
		v[i] *= multiplier
	}
}

func (v Vector) ReplaceVariable(index int, replacement Vector) {
	multiplier := v[index]
	for i, value := range replacement {
		if i == index {
			v[i] = 0
		} else {
			v[i] += value * multiplier
		}
	}
}

func (v Vector) String() string {
	return fmt.Sprint([]float64(v))
}
