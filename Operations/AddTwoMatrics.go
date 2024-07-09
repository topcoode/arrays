package operations

import (
	"github.com/gin-gonic/gin"
)

type ListNode struct {
	value int
	Next  *ListNode
}

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
*/
func Add(c *gin.Context) {
	l1 := ListNode{
		value: 2,
		Next: &ListNode{
			value: 4,
			Next: &ListNode{
				value: 3,
				Next:  nil,
			},
		},
	}
	l2 := ListNode{
		value: 5,
		Next: &ListNode{
			value: 6,
			Next: &ListNode{
				value: 4,
				Next:  nil,
			},
		},
	}

	addTwoNumbers(&l1, &l2)
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.value
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.value
			l2 = l2.Next
		}
		sum := carry + x + y
		carry = sum / 10
		current.Next = &ListNode{value: sum % 10}
		current = current.Next
	}
	if carry > 0 {
		current.Next = &ListNode{value: carry}
	}

	return dummy.Next
}
