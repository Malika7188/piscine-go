package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	ch := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = ch
		l.Tail = ch
	} else {
		l.Tail.Next = ch
		l.Tail = ch
	}
}
