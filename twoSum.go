package main

// brute
func twoSum_2(nums []int, target int) []int {
	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum_3(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, num := range nums {
		_, ok := hm[num]
		if ok {
			return []int{hm[num], i}
		}
		hm[target-num] = i

	}
	return nil
}

func twoSum(nums []int, target int) []int {
	mymap := make(map[int]int)
	for i := 0; i < len(nums); i++ {

		j, ok := mymap[target-nums[i]]
		if ok {
			return []int{j, i}
		}
		mymap[nums[i]] = i
	}
	return nil
}
