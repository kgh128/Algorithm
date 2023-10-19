/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 1. 재귀 호출
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    result := []int{root.Val}
    result = append(result, preorderTraversal(root.Left)...)
    result = append(result, preorderTraversal(root.Right)...)
    
    return result
}


// 2. 반복문
