package exercise

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number:", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	const diff = 0.0000001
	z := 1.0
	for math.Abs(x-z*z) > diff {
		z = z - (z*z-x)/(2*z)
	}

	return z, nil
}

func ErrorsMain() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
