// main.go
package main

import "fmt"

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(twoSum(nums, target)) // should print [0 1]
}