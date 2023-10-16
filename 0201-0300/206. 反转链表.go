package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	n1 := head
	n2 := head.Next.Next
	head = head.Next
	for head.Next != nil {

		n2 = head.Next
		head.Next = n1
		n1 = head
		head = n2
	}
	return head
}
