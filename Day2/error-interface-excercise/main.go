package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// Error ErrNegativeSqrt(-2).Error()
// returns "cannot Sqrt negative number: -2".
func (e ErrNegativeSqrt) Error() string {

}

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
