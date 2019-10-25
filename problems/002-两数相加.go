package main

import "fmt"

/*
002 两数相加 https://leetcode-cn.com/problems/add-two-numbers/

给出两个 非空 的链表用来表示两个非负的整数。
其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/
func main() {
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil,}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil,}}}
	ln := addTwoNumbers(&l1, &l2)
	for ln != nil {
		fmt.Println(ln.Val)
		ln = ln.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ln *ListNode
	n1 := l1
	n2 := l2
	tail := ln
	carry := 0
	for n1 != nil || n2 != nil || carry != 0 {
		v1 := 0
		v2 := 0
		if n1 != nil {
			v1 = n1.Val
		}
		if n2 != nil {
			v2 = n2.Val
		}
		sum := v1 + v2 + carry
		carry = sum / 10

		if ln == nil {
			ln = &ListNode{sum % 10, nil,}
			tail = ln
		} else {
			tail.Next = &ListNode{sum % 10, nil,}
			tail = tail.Next
		}

		if n1 != nil {
			n1 = n1.Next
		}
		if n2 != nil {
			n2 = n2.Next
		}
	}
	return ln
}

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/
