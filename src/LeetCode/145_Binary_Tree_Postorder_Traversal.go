/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 1. 재귀 함수로 풀기
// 시간복잡도: O(n), 공간복잡도: O(n)
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    result := postorderTraversal(root.Left)
    result = append(result, postorderTraversal(root.Right)...)

    return append(result, root.Val)
}


// 2. 반복문으로 풀기
// 시간복잡도: O(n), 공간복잡도: O(n)
// - 스택 이용 => 순회해야 하는 노드 리스트
// - 왼쪽 자식 노드 -> 오른쪽 자식 노드 -> 부모 노드 순서로 돌아야 한다.
// - 반복문으로 스택을 이용했을 때, 스택에서 꺼낸 노드의 값을 result 배열에 저장해야 한다.
// - 그런데, 스택에서 노드를 꺼내고 이후에 그 노드의 자식 노드를 스택에 집어넣는 매커니즘이다.
// - 따라서 필연적으로 자식 노드보다 부모 노드의 값이 먼저 result 배열에 저장될 수 밖에 없다.
// - 그래서 아예 result 배열의 끝에서 부터 값을 저장하기로 하였다. (정확히는 새로 추가되는 값은 result 배열의 가장 앞에 저장된다.)
// - 그러면 먼저 result 배열에 쓰인 값이 result 배열의 끝쪽에 위치한다.
// - 그러므로 스택에서 꺼낼 때도 원래 순회 순서와 반대로 노드를 꺼내야 한다. (부모 노드 -> 오른쪽 자식 노드 -> 왼쪽 자식 노드)
// - 스택은 LIFO이므로 더 먼저 방문해야 하는 노드를 더 나중에 넣어야 한다.
// - 부모 노드는 항상 자식 노드보다 먼저 방문하니 상관없고, 왼쪽 자식 노드를 먼저 스택에 넣고, 그 다음에 오른쪽 자식 노드를 스택에 넣는다.
// - 그럼으로써 오른쪽 자식 노드를 왼쪽 자식 노드보다 먼저 방문하도록 한다.
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    stack := []*TreeNode{root}
    result := []int{}

    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }

        result = append([]int{node.Val}, result...)
    }

    return result
}
