package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
