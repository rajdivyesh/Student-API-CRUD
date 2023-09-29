package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func insertAtEnd(head **Node, data int) {
	newNode := &Node{Data: data, Next: nil}
	if *head == nil {
		*head = newNode
	} else {
		current := *head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}
func printMyLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%d -->", current.Data)
		current = current.Next
	}
	fmt.Println("Nil")
}

func deleteItem(head **Node, target int) {
	if *head == nil {
		return
	}

	prev := *head
	current := prev.Next

	for current != nil {
		if current.Data == target {
			prev.Next = current.Next // Remove the target node
			return
		}
		prev = current
		current = current.Next
	}
}
func main4() {
	var head *Node
	insertAtEnd(&head, 1)
	insertAtEnd(&head, 2)
	insertAtEnd(&head, 3)
	insertAtEnd(&head, 4)
	fmt.Println("Voila, the linked list items are:")
	printMyLinkedList(head)

	var target int
	fmt.Print("Enter the value to delete: ")
	_, err := fmt.Scan(&target)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	deleteItem(&head, target)
	fmt.Println("After deleting target value:")
	printMyLinkedList(head)
}

/* the issue with previous setting was the method "insertAtEnd" function does not update the "head"
pointer outside the function because Golang uses pass-by-value for function arguments. Therefore,
to fix that problem we need to update the "head" pointer by passing a pointer to the head ("**Node")
to the "insertAtEnd" function.
SOME NOTES:
1. The "insertAtEnd" function now accepts a pointer to teh head of the linked list ("**Node")
so that it can modify the "head" pointer outside of the function
2. We used "*head" to access and update the "head" pointer within the function
*/
