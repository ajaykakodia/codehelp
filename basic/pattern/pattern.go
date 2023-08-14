package pattern

import (
	"fmt"
	"strconv"
)

func PatternPrint() {
	// Print Pascal Triangle
	pascalTriangle(6)
}

/*
Pascal Triangle -
1
1 1
1 2 1
1 3 3 1
1 4 6 4 1
1 5 10 10 5 1
1 6 15 20 15 6 1

formula - d=d*(row-col)/col [where row & col starts with 1 and d=1 at row level loop]
*/
func pascalTriangle(n int) {
	for r := 1; r <= n; r++ {
		d := 1
		for c := 1; c <= r; c++ {
			fmt.Print(strconv.Itoa(d) + " ")
			d = d * (r - c) / c
		}
		fmt.Println()
	}
}
