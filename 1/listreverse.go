package piscine

/*type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}*/

func ListReverse(l *List) {

	prev := l.Head
	prev = nil
	current := l.Head
	next := l.Head
	next = nil
	l.Tail = l.Head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
