package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

/*
WordCount counts words in s
*/
func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, word := range strings.Fields(s) {
		wc[word]++
	}
	fmt.Println(wc)
	return wc
}

func main() {
	wc.Test(WordCount)
}
