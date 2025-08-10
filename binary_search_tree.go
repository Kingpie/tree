package tree

import "fmt"

type BinarySearchTreeNode struct {
	val   int
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

func NewBinarySearchTreeNode(val int) *BinarySearchTreeNode {
	return &BinarySearchTreeNode{val: val}
}

// insert 向二叉搜索树中插入一个值，返回插入后的根节点
// root: 二叉搜索树的根节点指针
// val: 要插入的值
// 返回值: 插入新值后的二叉搜索树根节点指针
func insert(root *BinarySearchTreeNode, val int) *BinarySearchTreeNode {
	// 如果当前节点为空，创建新节点并返回
	if root == nil {
		return NewBinarySearchTreeNode(val)
	}

	// 根据二叉搜索树的性质递归插入
	// 如果插入值小于当前节点值，插入到左子树
	if val < root.val {
		root.left = insert(root.left, val)
	} else {
		// 否则插入到右子树
		root.right = insert(root.right, val)
	}

	// 返回当前节点作为根节点
	return root
}

func search(root *BinarySearchTreeNode, val int) bool {
	if root == nil {
		return false
	}
	if val < root.val {
		return search(root.left, val)
	} else if val > root.val {
		return search(root.right, val)
	}
	return true
}

// inOrder performs an in-order traversal of a binary search tree
// root: the root node of the binary search tree to traverse
func inOrder(root *BinarySearchTreeNode) {

	// Base case: if the current node is nil, return
	if root == nil {
		return
	}

	// Recursively traverse the left subtree
	inOrder(root.left)

	// Visit the current node by printing its value
	fmt.Println(root.val)

	// Recursively traverse the right subtree
	inOrder(root.right)
}

func preOrder(root *BinarySearchTreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	preOrder(root.left)
	preOrder(root.right)
}

// postOrder 对二叉搜索树进行后序遍历
// 后序遍历的顺序是：左子树 -> 右子树 -> 根节点
//
// 参数:
//
//	root - 二叉搜索树的根节点指针
//
// 返回值: 无
func postOrder(root *BinarySearchTreeNode) {
	// 递归终止条件：如果当前节点为空，则直接返回
	if root == nil {
		return
	}

	// 递归遍历左子树
	postOrder(root.left)

	// 递归遍历右子树
	postOrder(root.right)

	// 访问当前节点，输出节点值
	fmt.Println(root.val)
}

// levelOrder 对二叉搜索树进行层序遍历（广度优先遍历）
// 参数:
//
//	root - 二叉搜索树的根节点指针
//
// 返回值: 无
//
// 该函数使用队列实现层序遍历，从根节点开始，逐层从左到右访问所有节点，
// 并将每个节点的值打印到标准输出
func levelOrder(root *BinarySearchTreeNode) {
	// 空树检查
	if root == nil {
		return
	}

	// 初始化队列，将根节点入队
	queue := []*BinarySearchTreeNode{root}

	// 当队列不为空时，继续遍历
	for len(queue) > 0 {
		// 取出队列头部节点
		node := queue[0]
		queue = queue[1:]

		// 打印当前节点的值
		fmt.Println(node.val)

		// 将当前节点的左右子节点依次入队
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}
