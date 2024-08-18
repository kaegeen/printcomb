package main

import (
	"fmt"
	"strings"
)

func main() {
	printComb()
}

func printComb() {
	var result []string

	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				result = append(result, fmt.Sprintf("%d%d%d", i, j, k))
			}
		}
	}

	fmt.Println(strings.Join(result, ", "))
}
