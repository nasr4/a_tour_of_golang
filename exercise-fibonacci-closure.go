//Exercise: https://tour.golang.org/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	return func(x int) int {
		if x > 1 {
			minusOne, minusTwo := fibonacci(), fibonacci()
			
			return minusOne(x-1) + minusTwo(x-2)
		}
		
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
