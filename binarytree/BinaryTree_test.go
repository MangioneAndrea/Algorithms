package binarytree

import (
	"testing"
)

func TestExample1(t *testing.T) {
	root := &BinNode{
		Left:  &BinNode{},
		Right: &BinNode{},
	}
	_, isBalanced := root.IsBalanced()
	if !isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false")
	}
}

func TestExample2(t *testing.T) {
	root := &BinNode{
		Left: &BinNode{},
	}
	_, isBalanced := root.IsBalanced()
	if !isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false")
	}
}
func TestExample3(t *testing.T) {
	root := &BinNode{
		Left: &BinNode{
			Right: &BinNode{},
		},
	}
	balance, isBalanced := root.IsBalanced()
	if isBalanced {
		t.Errorf("Tree should not be balanced, but IsBalanced() returned true. Expected balance 2 got %d", balance)
	}
}
func TestExample4(t *testing.T) {
	root := &BinNode{
		Left: &BinNode{
			Right: &BinNode{},
		},
		Right: &BinNode{},
	}
	_, isBalanced := root.IsBalanced()
	if !isBalanced {
		t.Error("Tree should be balanced, but IsBalanced() returned false")
	}
}

func TestExample5(t *testing.T) {
	root := &BinNode{
		Left: &BinNode{
			Right: &BinNode{
				Right: &BinNode{},
			},
		},
	}
	balance, isBalanced := root.IsBalanced()
	if isBalanced {
		t.Errorf("Tree should not be balanced, but IsBalanced() returned true. Expected balance 3 got %d", balance)
	}
}
