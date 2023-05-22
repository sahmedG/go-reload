package proj

func ListAt(l *Node, pos int) *Node {
	if l == nil {
		return nil
	}

	cur := l
	i := 0

	for cur != nil && i < pos {
		cur = cur.Next
		i++
	}

	if i == pos {
		return cur
	}

	return nil
}
