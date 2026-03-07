// main.go
package main

import (
	"fmt"
	"flag"
)

func main() {

	algo := flag.String("algo", "brute-force", "the algorithm to use")
	flag.Parse()
	
	var nums = []int{2, 7, 11, 15}
	var target = 9
	
	switch *algo {
	case "hash":
		fmt.Println(twoSumHash(nums, target))
	default:
		fmt.Println(twoSumBrute(nums, target))
	}

}