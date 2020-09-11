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
	l1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	ln := addTwoNumbers(&l1, &l2)
	for ln != nil {
		fmt.Println(ln.Val)
		ln = ln.Next
	}
}

// 执行用时 4 ms 内存消耗 4.9 MB
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	tail := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		tail.Next = &ListNode{sum % 10, nil}
		tail = tail.Next
	}
	return head.Next
}

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/
