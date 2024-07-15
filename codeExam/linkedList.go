package codeExam

import "fmt"

type Node struct {
    data int
    next *Node
}

type LinkedList struct {
    head *Node
}
func (list *LinkedList) InsertFirst(data int) {
    newNode := &Node{data: data, next: list.head}
    list.head = newNode
}

func (list *LinkedList) InsertLast(data int) {
    newNode := &Node{data: data, next: nil}

    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

func (list *LinkedList) Display() {
    if list.head == nil {
        fmt.Println("Linked list is empty.")
        return
    }

    current := list.head
    for current != nil {
        fmt.Printf("%d ", current.data)
        current = current.next
    }
    fmt.Println()
}

func (list *LinkedList) Delete(data int) {
    if list.head == nil {
        fmt.Println("Linked list is empty. Nothing to delete.")
        return
    }

    if list.head.data == data {
        list.head = list.head.next
        return
    }

    current := list.head
    for current.next != nil {
        if current.next.data == data {
            current.next = current.next.next
            return
        }
        current = current.next
    }

    fmt.Printf("Element %d not found in the linked list.\n", data)
}

func (list *LinkedList) Search(data int) bool {
    if list.head == nil {
        return false
    }

    current := list.head
    for current != nil {
        if current.data == data {
            return true
        }
        current = current.next
    }

    return false
}

func (list *LinkedList) RemoveFirst() {
    if list.head == nil {
        fmt.Println("Linked list is empty. Nothing to remove.")
        return
    }

    list.head = list.head.next
}

func (list *LinkedList) RemoveLast() {
    if list.head == nil {
        fmt.Println("Linked list is empty. Nothing to remove.")
        return
    }

    if list.head.next == nil {
        list.head = nil
        return
    }

    current := list.head
    for current.next.next != nil {
        current = current.next
    }
    current.next = nil
}