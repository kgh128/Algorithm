/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/*
1. Hash Table 이용
- 시간복잡도: O(m+n), 공간복잡도: O(n)
- 키가 존재하는지를 빠르게 확인하기 위해서 map을 사용하는 것이므로 값은 아무거나 넣어도 된다.
- 더 많은 공간을 차지하는 int형(headA.Val)보다는 그냥 bool형을 값으로 넣는다.
=> 소비 시간과 공간이 줄어들었다.
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    mapA := make(map[*ListNode]bool)

    for headA != nil {
        mapA[headA] = true
        headA = headA.Next
    }

    for headB != nil {
        _, exists := mapA[headB]

        if exists {
            return headB
        }

        headB = headB.Next
    }

    return nil
}


/*
2. Two-Pointer 이용
- 시간복잡도: O(), 공간복잡도: O(1)
*/
