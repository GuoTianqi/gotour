package exercise

import (
	"fmt"
	"golang.org/x/tour/tree"
)

const QUICK = -1

func Walk(t *tree.Tree, ch chan int) {
	walkInner(t, ch)

	ch <- QUICK
}

func walkInner(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkInner(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walkInner(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1 := <-ch1
		v2 := <-ch2

		if v1 == v2 {
			if v1 == QUICK {
				return true
			}
		} else {
			return false
		}
	}
}

func EquivalentBinaryTreesMain() {
	fmt.Println("Same True: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Same False: ", Same(tree.New(1), tree.New(2)))
}
