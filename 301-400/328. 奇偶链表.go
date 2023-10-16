package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	root1 := head
	root2 := head.Next
	temp := head.Next

	for root1 != nil && root1.Next != nil {
		root1.Next = root1.Next.Next

		root1 = root1.Next
		root2.Next = root2.Next.Next
		root2 = root2.Next
	}

	root1.Next = temp

	return head

}
