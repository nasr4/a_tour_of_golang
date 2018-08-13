//Exercise: https://tour.golang.org/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var aCheck *int
	var a int
	if aCheck == nil {
		a := 0
		aCheck = &a
	}
	
	var results [10]int
	
	return func() int {	
		if a <= 1 {
			results[a] = a
		} else {
			results[a] = results[a-1] + results[a-2]
		}
		
		a++
		return results[a-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

