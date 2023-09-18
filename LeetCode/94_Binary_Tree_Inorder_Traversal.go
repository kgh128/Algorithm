/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 1. 반복문으로 풀기 - stack 이용
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    stack := []*TreeNode{root}
    result := []int{}

    for len(stack) > 0 {
        curr := stack[len(stack)-1]

        if curr.Left == nil {
            result = append(result, curr.Val)
            stack = stack[:len(stack)-1]

            if curr.Right != nil {
                stack = append(stack, curr.Right)
            }
        } else {
            stack = append(stack, curr.Left)
            curr.Left = nil // 왼쪽 자식 노드는 이미 stack에 넣었으므로 다음에 다시 왼쪽 자식 노드를 방문하지 않도록 표시한다.
        }
    }

    return result
}

// 2. 재귀함수로 풀기
// - inorder traversal(중위 순회): Left -> root -> Right 순서로 순회
// - 왼쪽으로 쭉 가다가 왼쪽 자식 노드가 없으면 그때 root로 올라오고, 그 다음에 오른쪽 자식으로 간다.
// - 순회 순서를 이용하여 append를 할 때 왼쪽 자식 트리 순회(inorderTraversal(root.Left)), root 노드의 값(root.Val), 오른쪽 자식 트리 순회(inorderTraversal(root.Right)) 순서대로 append한다.
// - 그리고 root가 nil이면 더이상 왼쪽 자식 노드가 존재하지 않는다는 의미이므로 return nil을 해서 자신의 부모 노드(이전 root)로 거슬러 올라간다.
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
}
