package main

import "fmt"

/*
025 K个一组翻转链表 https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
func main() {
	nodes := reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 4)
	for nodes != nil {
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

// 执行用时 4 ms 内存消耗 4.1 MB
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	if k > len(nodes) {
		return nodes[0]
	}
	ret := &ListNode{0, nil}
	tail := ret
	for i := 0; i < len(nodes)/k; i++ {
		for j := (i+1)*k - 1; j > i*k-1; j-- {
			tail.Next = nodes[j]
			tail = tail.Next
			tail.Next = nil
		}
	}
	r := len(nodes) % k
	if r != 0 {
		tail.Next = nodes[len(nodes)-r]
	}
	return ret.Next
}

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/
