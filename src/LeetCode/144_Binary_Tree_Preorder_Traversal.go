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
// - 스택 이용 => 순회해야 하는 노드 리스트
// - 현재 노드 -> 왼쪽 자식 노드 -> 오른쪽 자식 노드 순서로 돌아야 한다.
// - 현재 노드는 스택의 top에서 꺼내고, 현재 노드의 값은 결과 배열에 저장한다.
// - 현재 노드의 자식 노드들을 스택에 추가한다.
// - 스택은 LIFO이므로 더 먼저 방문해야 하는 노드를 더 나중에 넣어야 한다.
// - 따라서 오른쪽 자식 노드를 먼저 스택에 넣고, 그 다음에 왼쪽 자식 노드를 스택에 넣는다.
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    result := []int{}
    stack := []*TreeNode{root}

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        result = append(result, node.Val)

        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }

    return result
}
