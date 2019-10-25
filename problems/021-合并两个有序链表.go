package main

import "fmt"

/*
021 合并两个有序链表 https://leetcode-cn.com/problems/merge-two-sorted-lists/

*/
func main() {
	nodes := mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	for nodes != nil {
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

// 执行用时 4 ms 内存消耗 2.5 MB
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	var tail *ListNode
	if l1.Val < l2.Val {
		head = l1
		tail = l1
		l1 = l1.Next
	} else {
		head = l2
		tail = l2
		l2 = l2.Next
	}
	for {
		if l1 == nil && l2 == nil {
			break
		} else if l2 == nil {
			tail.Next = l1
			l1 = l1.Next
		} else if l1 == nil {
			tail.Next = l2
			l2 = l2.Next
		} else {
			if l1.Val < l2.Val {
				tail.Next = l1
				l1 = l1.Next
			} else {
				tail.Next = l2
				l2 = l2.Next
			}
		}
		tail = tail.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
