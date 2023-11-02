/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
1. 반복문 + 해시 테이블 사용
- 시간 복잡도: O(n), 공간 복잡도: O(n)
- 해시 테이블에 각 숫자의 개수를 저장
*/
func findMode(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    count := make(map[int]int)
    stack := []*TreeNode{root}

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if _, exists := count[node.Val]; exists {
            count[node.Val] += 1        
        } else {
            count[node.Val] = 1
        }

        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }

    maxCount := 0
    modes := []int{}

    for key := range count {
        if count[key] == maxCount {
            modes = append(modes, key)
        } else if count[key] > maxCount {
            maxCount = count[key]
            modes = []int{key}
        }
    }

    return modes
}
