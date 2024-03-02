package leetcode

func majorityElement2(nums []int) []int {
	elem1, elem2, count1, count2 := 0, 0, 0, 0
	for _, num := range nums {
		if count1 == 0 && num != elem2 {
			elem1 = num
			count1 = 1
		} else if count2 == 0 && num != elem1 {
			elem2 = num
			count2 = 1
		} else if elem1 == num {
			count1++
		} else if elem2 == num {
			count2++
		} else {
			count1--
			count2--
		}
	}
	target := len(nums)/3 + 1
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == elem1 {
			count1++
		} else if num == elem2 {
			count2++
		}
	}
	result := []int{}
	if count1 >= target {
		result = append(result, elem1)
	}
	if count2 >= target {
		result = append(result, elem2)
	}
	return result
}
