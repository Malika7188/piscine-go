package piscine

func ListPushFront(l *List, data interface{}) {
	ch := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = ch
		l.Tail = ch
	} else {
		l.Tail.Next = ch
		l.Tail = ch
	}
}
