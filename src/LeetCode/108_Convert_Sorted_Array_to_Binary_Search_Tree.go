/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    rootIdx := len(nums) / 2
    root := &TreeNode{Val: nums[rootIdx]}

    root.Left = sortedArrayToBST(nums[:rootIdx])
    root.Right = sortedArrayToBST(nums[rootIdx + 1:])

    return root
}
