package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	pos1 := 0

	for l != nil {
		if pos1 == pos {
			return l
		}
		l = l.Next
		pos1++
	}
	return nil
}
