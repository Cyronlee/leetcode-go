package main

import "fmt"

/*
024 两两交换链表中的节点 https://leetcode-cn.com/problems/swap-nodes-in-pairs/

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
*/
func main() {
	//nodes := swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}})
	nodes := swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	for nodes != nil {
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

// 执行用时 0 ms 内存消耗 2 MB
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ret *ListNode
	var tail *ListNode
	var temp *ListNode
	for head != nil {
		if head.Next != nil {
			temp = head.Next.Next
		} else {
			tail.Next = head
			return ret
		}
		if tail == nil {
			tail = head.Next
			ret = tail
		} else {
			tail.Next = head.Next
			tail = tail.Next
		}
		tail.Next = head
		tail = tail.Next
		// 防止循环链表
		tail.Next = nil
		head = temp
	}
	return ret
}

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/
