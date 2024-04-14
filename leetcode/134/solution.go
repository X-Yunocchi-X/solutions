package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	gas = append(gas, gas...)[:2*n-1]
	cost = append(cost, cost...)[:2*n-1]
	result := -1
	index := 0
	now := 0
	forward := 0
	for {
		if forward == n {
			result = index - forward
			break
		}
		if index == 2*n-1 {
			break
		}
		now += gas[index]
		if now < cost[index] {
			now = 0
			forward = 0
		} else {
			now -= cost[index]
			forward++
		}
		index++
	}

	return result
}
