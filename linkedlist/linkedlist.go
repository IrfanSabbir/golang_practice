package linkedListExample

import (
	"fmt"
	"log"
)

type node struct {
	next *node
	data int
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(val int) {
	newNode := node{data: val}

	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
}

func (l *LinkedList) printList() {
	if l.head == nil {
		return
	}
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
	log.Println()
	return
}

func (l *LinkedList) delteListitem(val int) {
	if l.head == nil {
		return
	}
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != val {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
	return
}

func (l *LinkedList) reverselist() {
	if l.head.next == nil {
		return
	}
	var prev *node
	cur := l.head

	for cur != nil {
		temp := cur.next
		cur.next = prev
		prev = cur
		cur = temp
	}
	l.head = prev
}

func LinkedListExample() {
	log.Println()
	fmt.Println("Linked list Examples.................\n")

	myList := LinkedList{}
	myList.prepend(20)
	myList.prepend(10)
	myList.prepend(25)
	myList.prepend(14)
	myList.prepend(24)
	myList.prepend(34)

	myList.printList()
	myList.delteListitem(14)
	myList.printList()
	// myList.reverselist()
	myList.reverselist()

	myList.printList()

	log.Println()
	return
}
