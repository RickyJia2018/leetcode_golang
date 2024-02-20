/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 */

// @lc code=start
type SingleNode struct{
	Val int
	Next *SingleNode
}

type MyLinkedList struct {
	DumHead *SingleNode
	Size int
}


func Constructor() MyLinkedList {
	dumHead := &SingleNode{
		Next: nil,

	}
	return MyLinkedList{
		DumHead: dumHead,
		Size: 0,
	}
}


func (this *MyLinkedList) Get(index int) int {
     if index<0||index>this.Size{
		return -1
	 }
	 cur := this.DumHead.Next
	 count:=0
	 for cur !=nil{
		if index == count{
			return cur.Val
		}
		cur = cur.Next
		count++
	 }
	 return -1
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &SingleNode{
		Val: val,
		Next: this.DumHead.Next,
	}
	this.DumHead.Next = newNode
	this.Size++
	this.PrintList()

}


func (this *MyLinkedList) AddAtTail(val int)  {
	newNode := &SingleNode{
		Val: val,
		Next: nil,
	}
	cur:= this.DumHead
	for cur.Next != nil{
		cur = cur.Next
	}
	cur.Next = newNode
	this.Size++
	this.PrintList()

}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index<0 || index>this.Size{
		return
	}
	count := 0
	cur := this.DumHead
	for cur.Next != nil{
		if count == index{
			break
		}
		cur = cur.Next
		count++

	}
		newNode := &SingleNode{
			Val: val,
			Next: cur.Next,
		}
		cur.Next = newNode
		this.Size++

	this.PrintList()

}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index < 0 || index >= this.Size {
		return
	}
    count:=0
	cur := this.DumHead
	for cur.Next!=nil{
		if count == index{
			cur.Next = cur.Next.Next
			this.Size --
			return
		}else{
			cur = cur.Next
			count++
		}
	}
	this.PrintList()
}

func (list *MyLinkedList) PrintList(){

	cur := list.DumHead
	for cur.Next != nil{
		fmt.Printf("%d ",cur.Next.Val)
		cur = cur.Next
	}
	fmt.Println("")
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

