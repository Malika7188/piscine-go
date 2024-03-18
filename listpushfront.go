package piscine

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
