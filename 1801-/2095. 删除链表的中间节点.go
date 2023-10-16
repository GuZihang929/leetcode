package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	root1 := head
	root2 := head

	for root2.Next != nil || root2.Next.Next != nil {
		root1 = root1.Next
		root2 = root2.Next.Next
	}

	root1.Next = root1.Next.Next
	return head

}
