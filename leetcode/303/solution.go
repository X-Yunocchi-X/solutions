package leetcode

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	arr := make([]int, len(nums)+1)
	for i := range nums {
		arr[i+1] = arr[i] + nums[i]
	}
	return NumArray{arr}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.arr[right+1] - this.arr[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
