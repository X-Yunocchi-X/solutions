package leetcode

type MyQueue struct {
	pushStack, popStack []int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}

func (this *MyQueue) Push(x int) {
	this.pushStack = append(this.pushStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.popStack) == 0 {
		for len(this.pushStack) != 0 {
			this.popStack = append(this.popStack, this.pushStack[len(this.pushStack)-1])
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
		}
	}
	result := this.popStack[len(this.popStack)-1]
	this.popStack = this.popStack[:len(this.popStack)-1]
	return result
}

func (this *MyQueue) Peek() int {
	if len(this.popStack) == 0 {
		for len(this.pushStack) != 0 {
			this.popStack = append(this.popStack, this.pushStack[len(this.pushStack)-1])
			this.pushStack = this.pushStack[:len(this.pushStack)-1]
		}
	}
	return this.popStack[len(this.popStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.pushStack) == 0 && len(this.popStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
