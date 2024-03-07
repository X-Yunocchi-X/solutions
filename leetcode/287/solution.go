package leetcode

func findDuplicate(nums []int) int {
	fast, slow := 0, 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		break
	}
	return 0
}
