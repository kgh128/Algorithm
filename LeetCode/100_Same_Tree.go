/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 1. 재귀함수로 풀기
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p != nil && q != nil && p.Val == q.Val {
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    } else if p == nil && q == nil {
        return true
    } else {
        return false
    }
}


// 2. 재귀함수로 풀기 - 더 짧은 버전
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p != nil && q != nil && p.Val == q.Val {
        return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
    } else {
        return p == q
    }
}
