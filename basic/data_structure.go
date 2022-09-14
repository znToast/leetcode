package basic

import (
	"strings"
)

// 数据结构

// 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
----剑指Offer 05 替换空格
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了30.59%的用户
*/
func ReplaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}

/*
----剑指Offer 06 从尾到头打印链表
执行用时：0 ms 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.6 MB 在所有 Go 提交中击败了100.00%的用户
*/
func ReversePrint(head *ListNode) []int {
	/* 	arr := []int{}
	   	耗时久 击败 5%
	   		dfs := func(node *ListNode) {}
	   	   	dfs = func(node *ListNode) {
	   	   		if node != nil {
	   	   			arr = append([]int{node.Val}, arr...)
	   	   			dfs(node.Next)
	   	   		}
	   	   	}
	   	   	dfs(head)
	   	return arr */
	count := 0 //击败 100%
	node := head
	for node != nil {
		count++
		node = node.Next
	}
	arr := make([]int, count)
	node = head
	for node != nil {
		count--
		arr[count] = node.Val
		node = node.Next
	}
	return arr
}

/*
----剑指Oferr 24 反转链表
执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了22.09%的用户
*/
func ReverseList(head *ListNode) *ListNode {
	var newHear *ListNode
	var index *ListNode
	node := head
	for node != nil {
		if newHear == nil {
			newHear = &ListNode{node.Val, nil}
		} else {
			index = newHear
			newHear = &ListNode{node.Val, index}
		}
		node = node.Next
	}
	return newHear
}

/*
----剑指Offer 58-II 左旋转字符串
*/
func ReverseLeftWords(s string, n int) string {
	n = n % len(s)
	s1 := s[:n]
	s2 := s[n:]
	return s2 + s1
}
