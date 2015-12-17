package example

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func MultipleResultsMain() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
