package lc

import "fmt"

// 反转单链表  反转只需要将curr 的上一个节点编程curr.next
func main() {
	head := &ListNode{
		data: 1, next: &ListNode{
			data: 2, next: nil,
		},
	}

	fmt.Println(reverseList(head))
}

// 迭代法
func reverseList(head *ListNode) *ListNode {
	var curr, prev *ListNode = head, nil
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
