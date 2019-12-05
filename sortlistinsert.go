package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	node := &NodeI{}
	node.Data = data_ref
	node.Next = nil

	if l == nil || l.Data >= node.Data {
		node.Next = l
		return node
	} else {
		a := l
		for a.Next != nil && a.Next.Data < node.Data {
			a = a.Next
		}
		node.Next = a.Next
		a.Next = node
	}
	return l
}
