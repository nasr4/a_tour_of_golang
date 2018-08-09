//Exercise: https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	eachWord := strings.Fields(s)
	wordCounter := make(map[string]int)
	
	for _, y := range eachWord {
		if _, ok := wordCounter[y]; ok {
			wordCounter[y] += 1
		} else {
			wordCounter[y] = 1
		}
	}
	
	return wordCounter
}

func main() {
	wc.Test(WordCount)
}

