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

func ListRemoveIf(l *List, dataRef interface{}) {
	var prev *NodeL
	current := l.Head

	for current != nil {
		if current.Data == dataRef {
			if prev == nil {
				l.Head = current.Next
			} else {
				prev.Next = current.Next
			}
			if current == l.Tail {
				l.Tail = prev
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
