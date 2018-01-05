package search

import (
	"math"

	"github.com/catorpilor/leetcode/utils"
)

func RecoverTree(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return utils.LevelOrderTravesal(root)
	}
	// inorder travesal
	// recursion so space is O(lgN) not constant
	var first, second *utils.TreeNode
	prev := &utils.TreeNode{Val: math.MinInt32}
	traverse(root, &first, &second, &prev)
	// fmt.Println(first, second)
	first.Val, second.Val = second.Val, first.Val
	return utils.LevelOrderTravesal(root)
}

func traverse(node *utils.TreeNode, first, second, prev **utils.TreeNode) {
	if node == nil {
		return
	}
	traverse(node.Left, first, second, prev)
	if (*first) == nil && (*prev).Val >= node.Val {
		*first = *prev
	}
	if (*first) != nil && (*prev).Val >= node.Val {
		*second = node
	}
	*prev = node
	traverse(node.Right, first, second, prev)
}
