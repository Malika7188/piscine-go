package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	char := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = char
		l.Tail = char
	} else {
		char.Next = l.Head
		l.Head = char
	}
}

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	count := 1
	for l.Head.Next != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
