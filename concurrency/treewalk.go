package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	go Walk(t1, ch1)
	ch2 := make(chan int, 10)
	go Walk(t2, ch2)

	for e1 := range ch1 {
		e2 := <-ch2
		if e1 != e2 {
			return false
		}
	}

	_, ok := <-ch2
	return !ok
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for e := range ch {
		fmt.Println(e)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
