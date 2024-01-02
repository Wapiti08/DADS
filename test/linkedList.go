package main

type SLLNode struct {
	next  *SLLNode
	value int
}

type Node interface {
	SetValue(v int)
	GetValue() int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return SNode.value
}

func NewSLLNode() *SLLNode {
	// new allocates memory space and return a pointer value
	return new(SLLNode)
}

type SingleLinkedList struct {
	head *SLLNode
	tail *SLLNode
}

func newSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.head == nil {
		// no head, then current node is the head of list
		list.head = newNode
	} else if list.tail == list.head {
		// there is only one node
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}
