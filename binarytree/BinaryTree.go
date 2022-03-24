package binarytree

import "math"

type BinNode struct {
	Value int
	Left  *BinNode
	Right *BinNode
}

func (root *BinNode) Rebalance() *BinNode {
	balance := root.GetBalance()
	if math.Abs(float64(balance)) < 1 {
		return root
	}
	var newRoot *BinNode
	var toMove *BinNode
	if balance > 0 {
		// rotateRight
		newRoot = root.Left
		toMove = newRoot.Right
		newRoot.Right = root
		root.Left = toMove
	} else {
		// rotateLeft
		newRoot = root.Right
		toMove = newRoot.Left
		newRoot.Left = root
		root.Right = toMove
	}
	return newRoot
}

func (root *BinNode) GetBalance() int {
	if root.Left == nil && root.Right == nil {
		return 0
	}

	if root.Left == nil {
		return -1 - int(math.Abs(float64(root.Right.GetBalance())))
	}
	if root.Right == nil {
		return 1 + int(math.Abs(float64(root.Left.GetBalance())))
	}
	return int(math.Abs(float64(root.Left.GetBalance()))) - int(math.Abs(float64(root.Right.GetBalance())))
}

func (root *BinNode) IsBalanced() (int, bool) {
	if root.Left == nil && root.Right == nil {
		return 1, true
	}

	if root.Left == nil {
		height, isBalanced := root.Right.IsBalanced()
		if !isBalanced || height > 1 {
			return height + 1, false
		}
		return height + 1, true
	}
	if root.Right == nil {
		height, isBalanced := root.Left.IsBalanced()
		if !isBalanced || height > 1 {
			return height + 1, false
		}
		return height + 1, true
	}
	height1, isBalanced := root.Left.IsBalanced()
	if !isBalanced {
		return height1 + 1, false
	}
	height2, isBalanced := root.Left.IsBalanced()
	if !isBalanced {
		return height2 + 1, false
	}
	return int(math.Max(float64(height1), float64(height2)) + 1), height2-height1 <= 1
}
