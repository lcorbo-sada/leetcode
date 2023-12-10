package main

import (
	"fmt"
)

func main() {
	input := []int{2, 7, 11, 15}
	target := 9
	output := twoSum(input, target)
	output2 := twoSum_2(input, target)
	output3 := twoSum_3(input, target)
	fmt.Println(output)  // [0,1]
	fmt.Println(output2) // [0,1]
	fmt.Println(output3) // [0,1]
}
