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
- 시간복잡도: O(n), 공간복잡도: O(h) - h는 트리의 높이
- dfs를 하므로 공간을 제일 많이 쓸 때는 h에 위치하는 leaf node에 갔을 때이다. 그때 h만큼 call stack이 호출되어 있다.
- 한 번 모든 트리를 dfs로 돌면서 계산한다.
- 사용자 정의 구조체에 현재 subtree의 sum, count, 최종 결과값을 담아서 전달한다.
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


/*
2. 포인터 변수 파라미터와 다중 반환값 + 재귀호출 이용
- 시간복잡도: O(n), 공간복잡도: O(h) - h는 트리의 높이
- dfs를 하므로 공간을 제일 많이 쓸 때는 h에 위치하는 leaf node에 갔을 때이다. 그때 h만큼 call stack이 호출되어 있다.
- 한 번 모든 트리를 dfs로 돌면서 계산한다.
- golang에서는 함수의 다중 반환값을 지원한다. 현재 subtree의 sum과 count를 모두 반환한다.
- golang에서도 함수의 파라미터로 포인터 변수를 넣을 수 있디. 최종 결과를 저장하는 result 변수의 포인터를 전달하면서, 각 함수 호출이 모두 같은 result 변수에 값을 쓸 수 있도록 한다.
=> result를 전역 변수로 선언하고 해당 result 변수를 모든 함수 호출이 사용하는 것과 같은 효과
*/
func traverseSubtree(root *TreeNode, result *int) (int, int) {
    if root == nil {
        return 0, 0
    }

    leftSum, leftCount := traverseSubtree(root.Left, result)
    rightSum, rightCount := traverseSubtree(root.Right, result)

    sum := root.Val + leftSum + rightSum
    count := 1 + leftCount + rightCount

    if root.Val == (sum / count) {
        *result++
    }

    return sum, count
}

func averageOfSubtree(root *TreeNode) int {
    result := 0
    traverseSubtree(root, &result)
    return result
}
