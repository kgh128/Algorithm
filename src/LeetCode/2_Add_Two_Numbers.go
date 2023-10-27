/*
<두 링크드 리스트에 저장된 수를 더하기>
- l1, l2는 링크드 리스트의 head 포인터 값이기 때문에 len(l1), len(l2)로 길이를 알아낼 수 없다.
=> 둘 다 nil 값이 될 때까지 반복문을 돈다.
=> 더 긴 길이의 링크드 리스트의 값을 모두 더할 때까지 반복문을 돌게 된다.

- 결과 배열도 링크드 리스트로 리턴해야 한다.
=> 빈 링크드 리스트 노드 만들기: node := &ListNode{}
=> 값을 넣은 링크드 리스트 노드 만들기: node := &ListNode{0, nil}
=> 다음 노드로 이동하기: node := node.Next
=> 현재 노드에 값 집어넣기: node.Val = 0
=> 현재 노드에 다음 노드 주소 넣기: node.Next = &ListNode{} -> 이거는 예시일 뿐 노드의 주소가 들어가기만 하면 된다.
*/


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 1. 결과 링크드 리스트
    result := &ListNode{}  // head 노드
    curNode := result      // 현재 노드

    // 2. 더 긴 링크드 리스트의 원소를 다 더할 때까지 반복문 돌기
    // 링크드 리스트가 끝나는 조건: 현재 노드의 포인터 값이 nil
    for l1 != nil || l2 != nil {
        // 현재 노드가 nil이 아니라면 그 값을 curNode.Val에 더하기(결과 링크드 리스트의 현재 노드)
        // 그리고 다음 노드로 넘어가기
        // 1) l1과 l2 모두 nil이 아님 => 둘의 값이 모두 더해짐
        // 2) 둘 중 하나가 nil => 하나의 값만 더해짐. (nil인 것은 0으로 취급되는 것과 같은 효과를 가짐.)
        // 3) 둘 다 nil => 아무 값도 안 더해짐. (curNode는 그냥 빈 노드)
        if l1 != nil {
            curNode.Val += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            curNode.Val += l2.Val
            l2 = l2.Next
        }

        // 3. curNode.Val가 두 자리 수인 경우
        // curNode.Val에서 10을 빼서 일의 자릿수만 남긴다.
        // 다음 노드의 Val에 1을 넣고 해당 노드를 curNode.Next로 설정한다.
        // => 만약 l1과 l2가 모두 nil이어서 다음에 더 계산할 것이 없더라도 넘어가는 십의 자릿수 1이 있기 때문에
        //    결과 링크드 리스트 마지막에 하나의 노드가 더 필요하다.
      
        // 4. curNode.Val이 한 자리 수인데, l1과 l2 중 하나라도 nil이 아니어서 계산할 것이 남아있는 경우
        // curNode.Next에 그냥 빈 ListNode의 포인터 값을 넣는다.

        // +) 다음으로 넘어가는 십의 자릿수도 없고, l1과 l2가 둘 다 nil이라면 
        //    더이상 계산할 것이 없으므로 결과 링크드 리스트에 더 노드를 추가할 필요가 없다.
        // => 현재 l1과 l2는 이번 반복문에서 계산에 사용한 노드가 아니라 그 다음 노드이다.
        // => 즉, 현재 l1과 l2가 모두 nil이라면 그 다음 계산이 없다는 의미이고, 하나라도 nil이 아니라면 그 다음 계산이 존재한다는 의미이다.
        // => 다음 계산이 남아있는 경우만 결과 링크드 리스트에 노드를 추가하므로 가장 마지막에 쓸데없는 노드가 추가되지 않는다.
        if curNode.Val > 9 {
            curNode.Val -= 10
            curNode.Next = &ListNode{1, nil}
        } else if l1 != nil || l2 != nil {
            curNode.Next = &ListNode{}
        }

        // 5. 결과 링크드 리스트의 노드도 다음 노드로 넘어간다.
        curNode = curNode.Next
    }

    // 6. 결과 링크드 리스트 리턴
    return result
}
