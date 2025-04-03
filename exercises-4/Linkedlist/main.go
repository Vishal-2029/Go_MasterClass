package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}



func (s *SinglyLinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if s.head == nil {
		s.head = newNode
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}


func (s *SinglyLinkedList) Display() {
	current := s.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

type DllNode struct {
	data int
	prev *DllNode
	next *DllNode
}

type DoublyLinkedList struct {
	head *DllNode
}

func (d *DoublyLinkedList) Insert(data int) {
	newNode := &DllNode{data: data}
	if d.head == nil {
		d.head = newNode
		return
	}
	current := d.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
}


func (d *DoublyLinkedList) DisplayForward(){
	current := d.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (d *DoublyLinkedList) DisplayBackward() {
	current := d.head
	if current == nil {
		fmt.Println("List is empty")
		return
	}

	for current.next != nil {
		current = current.next
	}
	
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}


func main() {
	fmt.Println("Singly Linked List:")
	s := &SinglyLinkedList{}
	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Display()

	fmt.Println("Doubly Linked List:")
	d := &DoublyLinkedList{}
	d.Insert(1)
	d.Insert(2)
	d.Insert(3)
	d.DisplayForward()
	d.DisplayBackward()
}
