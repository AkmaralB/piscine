package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}

// func ListPushBack(l *List, data interface{}) {
// 	n := &NodeL{Data: data}
// 	current := l.Head
// 	if l.Head == nil {
// 		l.Head = n
// 		return
// 	} else {
// 		for current.Next != nil {
// 			current = current.Next
// 		}
// 		current.Next = n
// 	}

// }
