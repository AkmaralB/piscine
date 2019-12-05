package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	text := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = text
	} else {
		text.Next = l.Head
		l.Head = text
	}
}
