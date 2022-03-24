package binarytree

import (
	"testing"
)

func TestExample1(t *testing.T) {
	root := &BinNode{
		Value: 2,
		Left:  &BinNode{Value: 1},
		Right: &BinNode{Value: 3},
	}
	_, isBalanced := root.IsBalanced()
	if !isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false")
	}
}

func TestExample2(t *testing.T) {
	root := &BinNode{
		Value: 1,
		Left:  &BinNode{Value: 0},
	}
	_, isBalanced := root.IsBalanced()
	if !isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false")
	}
}
func TestExample3(t *testing.T) {
	root := &BinNode{
		Value: 2,
		Left: &BinNode{
			Value: 0,
			Right: &BinNode{Value: 1},
		},
	}
	balance, isBalanced := root.IsBalanced()
	if isBalanced {
		t.Errorf("Tree should not be balanced, but IsBalanced() returned true. Expected balance 2 got %d", balance)
	}
	newRoot:=root.Rebalance()
	balance, isBalanced = newRoot.IsBalanced()
	if isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false", balance)
	}

}
func TestExample4(t *testing.T) {
	root := &BinNode{
		Value: 2,
		Left: &BinNode{
			Value: 0,
			Right: &BinNode{Value: 1},
		},
		Right: &BinNode{Value: 3},
	}
	_, isBalanced := root.IsBalanced()
	if !isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false")
	}
}

func TestExample5(t *testing.T) {
	root := &BinNode{
		Value: 4,
		Left: &BinNode{
			Value: 1,
			Right: &BinNode{
				Value: 2,
				Right: &BinNode{Value: 3},
			},
		},
	}
	balance, isBalanced := root.IsBalanced()
	if isBalanced {
		t.Errorf("Tree should not be balanced, but IsBalanced() returned true. Expected balance 3 got %d", balance)
	}
	newRoot:=root.Rebalance()
	balance, isBalanced = newRoot.IsBalanced()
	if isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false", balance)
	}
}
