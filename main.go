package main

import (
	"fmt"
	"slices"
	"strings"
)

// 旋转字符串
func rotateString(s string, goal string) bool {
    return len(s) == len(goal) && strings.Contains(s + s, goal)
}

// 二叉树的最大深度
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Left.Right)) + 1
}

// 相同的二叉树
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if (p == nil && q == nil) {
		return true
	}
	if (p != nil && q != nil && p.Val == q.Val) {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
    return checkTree(root.Left, root.Right)
}

func checkTree(p,q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && checkTree(p.Left, q.Right) && checkTree(p.Right, q.Left)
}

// 旋转数组
func rotate(nums []int, k int)  {
    k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

func main() {
	fmt.Println(rotateString("adcd", "cad"))
}