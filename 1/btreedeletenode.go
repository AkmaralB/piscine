package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Data < root.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if node.Data > root.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		if root.Left == nil {
			a := root.Right
			root = nil
			return a
		} else if root.Right == nil {
			a := root.Left
			root = nil
			return a
		}
		a := BTreeMin(root.Right)

		root.Data = a.Data
		root.Right = BTreeDeleteNode(root.Right, a)
	}
	return root
}
