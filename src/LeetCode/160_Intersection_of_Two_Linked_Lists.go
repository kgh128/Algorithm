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
- 시간복잡도: O(m+n), 공간복잡도: O(1)
- 이 문제에서 쟁점은 교집합 나오기 전의 링크드 리스트 A와 B의 노드의 개수가 다르다는 것이다.
- 그래서 하나의 포인터가 하나의 링크드 리스트만 한 번 도는 것으로는 문제를 해결할 수 없다.
- 따라서 ptr1과 ptr2가 각각 headA와 headB로 시작하고, 각 링크드 리스트를 다 돌았을 경우에는 자신이 돌지 않은 링크드 리스트를 처음부터 돈다.
- 이 와중에 ptr1과 ptr2가 같으면 교집합 노드를 찾은 것이므로 해당 노드를 리턴한다.
ex) A = a1 -> a2 -> c1 -> c2
    B = b1 -> b2 -> b3 -> c1 -> c2
    ptr1의 순회: a1 -> a2 -> c1 -> c2 -> b1 -> b2 -> b3 -> c1 -> c2
    ptr2의 순회: b1 -> b2 -> b3 -> c1 -> c2 -> a1 -> a2 -> c1 -> c2
    이처럼 처음 c1에 도달했을 때는 서로의 시점이 다르지만, 두번째 c1에 도달했을 때는 서로의 시점이 같다.
- A의 교집합 전의 노드 개수를 a, B의 교집합 전의 노드 개수를 b, 교집합 노드 개수를 c라고 하면
- 두 포인터 모두 (a + b + c)의 개수만큼 루프를 돌고 같은 시점에 교집합 노드에 도달한다.
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    ptr1 := headA
    ptr2 := headB

    for ptr1 != ptr2 {
        if ptr1 == nil {
            ptr1 = headB
        } else {
            ptr1 = ptr1.Next
        }

        if ptr2 == nil {
            ptr2 = headA
        } else {
            ptr2 = ptr2.Next
        }
    }

    return ptr1
}
