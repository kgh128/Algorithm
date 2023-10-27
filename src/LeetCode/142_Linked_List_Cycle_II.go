/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/*
2. Two-Pointer 이용 (Floyd's Cycle-Finding Algorithm)
- 시간복잡도: O(n), 공간복잡도: O(1) -> 포인터 2개의 공간만 필요
- 플로이드의 순환 찾기 알고리즘은 빠른 포인터가 느린 포인터를 따라잡는다면 cycle이 존재한다는 알고리즘이다.
- 빠른 포인터: 2칸씩 이동
- 느린 포인터: 1칸씩 이동
- 빠른 포인터의 값이 느린 포인터의 값과 같아진다면 빠른 포인터가 cycle 한 바퀴를 돌아서 따라잡은 것이다.
- 빠른 포인터가 느린 포인터를 만나지 못하고 linked list의 끝을 만난다면 cycle이 존재하지 않는 것이다.
- 빠른 포인터와 느린 포인터가 만났을 때 느린 포인터를 시작점으로 보내고 두 포인터 모두 1칸씩 이동한다.
- 그러다가 둘이 처음으로 만나는 지점이 cycle의 시작점이다.
- 출처: https://fierycoding.tistory.com/45
*/
func detectCycle(head *ListNode) *ListNode {
    slowPtr, fastPtr := head, head

    for fastPtr != nil && fastPtr.Next != nil {
        slowPtr = slowPtr.Next
        fastPtr = fastPtr.Next.Next

        if slowPtr == fastPtr {
            slowPtr = head

            for slowPtr != fastPtr{
                slowPtr = slowPtr.Next
                fastPtr = fastPtr.Next
            }

            return slowPtr
        }
    }

    return nil
}
