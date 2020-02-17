package main

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func ExampleWalk() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for e := range ch {
		fmt.Println(e)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
func TestSame(t *testing.T) {
	if Same(tree.New(1), tree.New(1)) != true {
		t.Errorf("Should be same")
	}

	if Same(tree.New(1), tree.New(2)) != false {
		t.Errorf("Should not be same")
	}
}
