package piscine

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
