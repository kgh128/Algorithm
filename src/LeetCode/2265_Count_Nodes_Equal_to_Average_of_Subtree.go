/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
1. 사용자 정의 구조체 + 재귀 호출 이용
- 시간복잡도: O(n)
*/
type SubtreeInfo struct {
    TotalCount int
    TotalSum int
    GoalCount int
}

func subtreeTraversal(root *TreeNode) SubtreeInfo {
    if root == nil {
        return SubtreeInfo{TotalCount: 0, TotalSum: 0, GoalCount: 0}
    }

    left := subtreeTraversal(root.Left)
    right := subtreeTraversal(root.Right)

    result := SubtreeInfo{}
    result.TotalCount = 1 + left.TotalCount + right.TotalCount
    result.TotalSum = root.Val + left.TotalSum + right.TotalSum
    result.GoalCount = left.GoalCount + right.GoalCount

    if root.Val == (result.TotalSum / result.TotalCount) {
        result.GoalCount++
    }

    return result
}

func averageOfSubtree(root *TreeNode) int {
    result := subtreeTraversal(root)
    return result.GoalCount
}
