/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
    main []int
	backUp []int
}


func Constructor() MyStack {
    return MyStack{
		main: make([]int,0),
		backUp: make([]int,0),
	}
}


func (this *MyStack) Push(x int)  {
    this.main = append(this.main, x)
}


func (this *MyStack) Pop() int {
   if len(this.main) > 0 {
	res := this.main[len(this.main) - 1]
	this.main = this.main[:len(this.main) - 1]
	return res
   }
	return -1
}


func (this *MyStack) Top() int {
    return this.main[len(this.main) - 1]
}


func (this *MyStack) Empty() bool {
    return len(this.main) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

