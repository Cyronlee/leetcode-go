package main

import "fmt"

/*
023 合并K个排序链表 https://leetcode-cn.com/problems/merge-k-sorted-lists/

合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/
func main() {
	list := []*ListNode{
		&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		&ListNode{4, &ListNode{5, &ListNode{6, nil}}},
		&ListNode{7, &ListNode{8, &ListNode{9, nil}}},
	}
	nodes := mergeKLists(list)
	for nodes != nil {
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

// 执行用时 328 ms 内存消耗 5.3 MB
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{0, nil}
	tail := head
	for {
		k := -1
		min := -999
		// 找出val最小的第K个元素
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if min == -999 || lists[i].Val < min {
				min = lists[i].Val
				k = i
			}
		}
		if k == -1 {
			return head.Next
		}
		tail.Next = lists[k]
		tail = tail.Next
		lists[k] = lists[k].Next
	}
}

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/
