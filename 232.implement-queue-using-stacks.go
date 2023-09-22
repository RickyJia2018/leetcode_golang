/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	stackIn []int
	stackOut []int
}


func Constructor() MyQueue {
   return MyQueue{
	stackIn: make([]int,0),
	stackOut: make([]int,0),
   }

}


func (this *MyQueue) Push(x int)  {
    this.stackIn = append(this.stackIn, x)
}


func (this *MyQueue) Pop() int {
    
	if len(this.stackOut) == 0 {
		for len(this.stackIn) >0 {
			this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
			this.stackIn = this.stackIn[0:len(this.stackIn)-1]
		}
	}
	if len(this.stackOut) == 0 {
        return -1
    }
	result := this.stackOut[len(this.stackOut)-1]
	this.stackOut = this.stackOut[0:len(this.stackOut)-1]
	return result
}


func (this *MyQueue) Peek() int {
	val := this.Pop()
	if val == -1 {
        return -1
    }
	this.stackOut = append(this.stackOut, val)
    return val
}


func (this *MyQueue) Empty() bool {
    return len(this.stackIn) == 0 && len(this.stackOut) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

