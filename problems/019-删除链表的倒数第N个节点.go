package main

import "fmt"

/*
019 删除链表的倒数第N个节点 https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

*/
func main() {
	head := removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 2)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// 执行用时 0 ms 内存消耗 2.2 MB
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	var nodes []*ListNode
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	if n == len(nodes) {
		return nodes[1]
	}
	if n == 1 {
		nodes[len(nodes)-2].Next = nil
	} else {
		nodes[len(nodes)-n-1].Next = nodes[len(nodes)-n+1]
	}
	return nodes[0]
}

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/
