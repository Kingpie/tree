package tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := NewBinarySearchTreeNode(6)
	insert(bst, 1)
	insert(bst, 3)
	insert(bst, 2)
	insert(bst, 7)
	insert(bst, 9)

	if search(bst, 4) {
		t.Error("Search failed")
	}

	if !search(bst, 9) {
		t.Error("Search failed")
	}

	fmt.Println("Inorder traversal:")
	inOrder(bst)
	fmt.Println("Preorder traversal:")
	preOrder(bst)
	fmt.Println("Postorder traversal:")
	postOrder(bst)
	fmt.Println("Levelorder traversal:")
	levelOrder(bst)
}
