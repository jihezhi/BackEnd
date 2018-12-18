package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

const epsilon = 1e-4

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x <= 0 {
		return x, ErrNegativeSqrt(x)
	}
	tempFunc := func(z float64) float64 { return z - (z*z-x)/(2*z) }
	for i, z2 := 0, tempFunc(z); i < 10; i, z, z2 = i+1, z2, tempFunc(z2) {
		if math.Abs(z-z2) < epsilon {
			break
		}
	}
	return z, nil
}
