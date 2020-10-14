package leetcode

import (
	"fmt"
	"time"
)

func strtoint(str string) int {
	if len(str) == 0 {
		return 0
	}
	var maxValue = 2147483647
	var minValue = -2147483648

	var tmp = 1
	var i int
	//跳过空格
	for i < len(str) && str[i] == ' ' {
		i = i + 1
	}
	if i == len(str) {
		return 0
	}
	//跳过符号
	if str[i] == '-' || str[i] == '+' {
		if str[i] == '-' {
			tmp = -1
		}
		i = i + 1
	}
	var num int
	for i < len(str) {
		//如果不是
		if str[i] < '0' || str[i] > '9' {
			break
		}
		//累计数字
		num = num*10 + int(str[i]-'0')
		//越界
		if num > maxValue {
			if tmp == 1 {
				return maxValue
			} else {
				return minValue
			}
		}
		i = i + 1
	}
	return tmp * num
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next //移动指针
	}
}

func outputInt(p *ListNode) []int { //遍历
	var output, res []int
	for p != nil {
		//fmt.Println(*p)
		output = append(output, p.Val)
		p = p.Next //移动指针
	}
	n := len(output)
	for i := n - 1; i >= 0; i-- {
		res = append(res, output[i])
	}
	return res
}

//尾插法
func creatListNode(input []int) *ListNode {
	var output = new(ListNode)
	output.Val = input[0]
	var tail *ListNode
	tail = output
	n := len(input)
	for i := 1; i < n; i++ {
		var tmp = ListNode{Val: input[i]}
		(*tail).Next = &tmp
		tail = &tmp
	}
	return output
}

type Node struct {
	data int
	next *Node
}

func Shownode1(p *Node) { //遍历
	for p != nil {
		fmt.Println(p.data)
		p = p.next //移动指针
	}
}
func test111() {
	var head = new(Node)
	head.data = 0
	var tail *Node
	tail = head //tail用于记录最末尾的结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		(*tail).next = &node
		tail = &node
	}
	Shownode1(head) //遍历结果
}

var POOL = 100

func groutine1(p chan int) {

	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("groutine-1:", i)
		}
	}
}

func groutine2(p chan int) {

	for i := 1; i <= POOL; i++ {
		<-p
		// p <- i
		if i%2 == 0 {
			fmt.Println("groutine-2:", i)
		}
	}
}

func Test1013() {
	msg := make(chan int)

	go groutine1(msg)
	go groutine2(msg)

	time.Sleep(time.Second * 1)

}

func chantest() {
	var NUM int
	NUM = 10000
	chanTest := make(chan int, 10)
	go func() {
		for i := 1; i <= NUM; i++ {
			chanTest <- i
			if i%2 == 1 {
				fmt.Println("奇数:", i)
			}
		}
	}()
	go func() {
		for i := 1; i <= NUM; i++ {
			<-chanTest
			if i%2 == 0 {
				fmt.Println("偶数:", i)
			}
		}
	}()
	time.Sleep(time.Second * 2)
}
