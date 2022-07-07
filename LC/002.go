package lc

import "fmt"

/**
两数相加，链表存储的是逆序 数字的数，同时遍历两个链表 并求每一位的和  记录进位数carry
链表长度不一致 如何遍历  结束 是两个都为nil
*/
type ListNode struct {
	data int
	next *ListNode
	prev *ListNode
}

func main() {
	var nun1 ListNode = ListNode{1, nil, nil}
	var nun2 ListNode = ListNode{2, nil, nil}

	sum1 := &ListNode{
		1, &nun1, nil,
	}
	sum := &ListNode{
		2, &nun2, nil,
	}
	a := addTwoNumbers(sum1, sum)
	for a != nil {
		fmt.Println(a.data)
		a = a.next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail, head *ListNode

	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.data
			l1 = l1.next
		}
		if l2 != nil {
			n2 = l2.data
			l2 = l2.next
		}
		sum := n1 + n2 + carry
		// 数字与进位数
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{data: sum}
			tail = head
		} else {
			tail.next = &ListNode{data: sum}
			tail = tail.next
		}
	}

	return head
}
