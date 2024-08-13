package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if node == nil {
		return root
	} else if node.Left == nil && node.Right == nil {
		if node == root {
			return nil
		}
		if node == node.Parent.Right {
			node.Parent.Right = nil
		} else {
			node.Parent.Left = nil
		}
		return root
	} else if node.Left == nil {
		if node == root {
			node.Right.Parent = nil
			return node.Right
		}
		if node == node.Parent.Right {
			node.Parent.Right = node.Right
		} else {
			node.Parent.Left = node.Right
		}
		node.Right.Parent = node.Parent
		return root
	} else if node.Right == nil {
		if node == root {
			node.Left.Parent = nil
			return node.Left
		}
		if node == node.Parent.Right {
			node.Parent.Right = node.Left
		} else {
			node.Parent.Left = node.Left
		}
		node.Left.Parent = node.Parent
		return root
	}
	nodeSuccessor := node.Right
	for nodeSuccessor.Left != nil {
		nodeSuccessor = nodeSuccessor.Left
	}
	if nodeSuccessor == node.Right {
		node.Right = nodeSuccessor.Right
		if nodeSuccessor.Right != nil {
			nodeSuccessor.Right.Parent = node
		}
	} else {
		nodeSuccessor.Parent.Left = nodeSuccessor.Right
		if nodeSuccessor.Right != nil {
			nodeSuccessor.Right.Parent = nodeSuccessor.Parent
		}
	}
	nodeSuccessor.Left = node.Left
	nodeSuccessor.Right = node.Right
	node.Left.Parent = nodeSuccessor
	node.Right.Parent = nodeSuccessor
	nodeSuccessor.Parent = node.Parent
	if node == root {
		return nodeSuccessor
	}
	if node == node.Parent.Left {
		node.Parent.Left = nodeSuccessor
	} else {
		node.Parent.Right = nodeSuccessor
	}
	return root
}
