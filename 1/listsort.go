package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	cur := l
	if cur == nil {
		return nil
	}
	cur.Next = ListSort(cur.Next)

	if cur.Next != nil && cur.Data > cur.Next.Data {
		cur = reverse(cur)
	}
	return cur
}

func reverse(l *NodeI) *NodeI {
	prev := l
	next := l.Next
	numb := next
	for next != nil && l.Data > next.Data {
		prev = next
		next = next.Next
	}
	prev.Next = l
	l.Next = next
	return numb
}
