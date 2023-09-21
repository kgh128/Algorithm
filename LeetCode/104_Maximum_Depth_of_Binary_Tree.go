/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 1. DFS - 재귀 호출
// 시간: 0ms, 공간: 4.2MB
func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
