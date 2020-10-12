package leetcode

import (
	"fmt"
	"testing"
)

func Test_strtoint(t *testing.T) {
	//空格符号在前
	int1 := strtoint("     -42")
	fmt.Println(int1)
	int1 = strtoint("42")
	fmt.Println(int1)
	//英文空格在后
	int1 = strtoint("4193 with words")
	fmt.Println(int1)
	//英文空格在前
	int1 = strtoint("ords and 987")
	fmt.Println(int1)
	//越界
	int1 = strtoint("-91283472332")
	fmt.Println(int1)
	//learndemo()
}

func Test_addTwoNumbers(t *testing.T) {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 2
	head.Next = node1
	var node2 = new(ListNode)
	node2.Val = 3
	node1.Next = node2

	var head2 = new(ListNode)
	head2.Val = 4
	var node21 = new(ListNode)
	node21.Val = 5
	head2.Next = node21
	var node22 = new(ListNode)
	node22.Val = 6
	node21.Next = node22

	// abc := addTwoNumbers(head, head2)
	Shownode(head)
	Shownode(head2)
	a := outputInt(head)
	b := outputInt(head2)
	fmt.Println(a)
	fmt.Println(b)
	newhead := creatListNode(a)
	Shownode(newhead)
	//	test111()
}
