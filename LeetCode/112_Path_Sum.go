/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 1. 원래 내 코드: 시간 - 0ms
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    if root.Left == nil && root.Right == nil {
        if root.Val == targetSum {
            return true
        } else {
            return false
        }
    }

    targetSum -= root.Val
    return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}


// 2. leaf node의 값과 targetSum과의 비교를 간략화: - 시간 6ms
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }

    targetSum -= root.Val

    if root.Left == nil && root.Right == nil {
        return targetSum == 0
    }

    return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
