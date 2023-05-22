package proj

func ListSize(l *List) int {
	counter := 0
	current := l.Head

	for current != nil {
		counter++
		current = current.Next
	}
	return counter
}
