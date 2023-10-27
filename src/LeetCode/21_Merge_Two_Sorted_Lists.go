/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 1. for문을 두 list가 모두 빌 때까지 돌기
// +) 현재 노드의 Val에 list1.Val이나 list2.Val을 넣기
// 남은 리스트가 하나밖에 없으면 for문 돌면서 남은 리스트 값을 하나씩 mergedList에 넣기
// 시간: 0ms
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  if list1 == nil && list2 == nil {
    return nil
  }

  mergedList := &ListNode{}
  curNode := mergedList

    for list1 != nil || list2 != nil {
      if list1 == nil {
        curNode.Val = list2.Val
        list2 = list2.Next
      } else if list2 == nil {
        curNode.Val = list1.Val
        list1 = list1.Next
      } else {
        if list1.Val < list2.Val {
          curNode.Val = list1.Val
          list1 = list1.Next
        } else {
          curNode.Val = list2.Val
          list2 = list2.Next
        }
      }

      if list1 != nil || list2 != nil {
        curNode.Next = &ListNode{}
      }
      curNode = curNode.Next
    }

    return mergedList
}

// 2. for문을 두 list 중 하나라도 비면 그만 돌기
// +) 현재 노드의 Next에 list1이나 list2의 현재 노드의 주소값을 넣기
// for문을 빠져나가면 결국 남은 리스트는 하나밖에 없으니 그 리스트를 mergedList에 갖다붙이기
// 시간: 3ms
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  mergedList := &ListNode{}
  curNode := mergedList

    for list1 != nil && list2 != nil {
      if list1.Val < list2.Val {
          curNode.Next = list1
          list1 = list1.Next
      } else {
          curNode.Next = list2
          list2 = list2.Next
      }

      curNode = curNode.Next
    }

    if list1 == nil {
      curNode.Next = list2
    } else {
      curNode.Next = list1
    }

    return mergedList.Next
}

// 3. 재귀함수로 풀기
// 시간: 3ms
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
  if list1 == nil {
    return list2
  }

  if list2 == nil {
    return list1
  }

  if list1.Val < list2.Val {
    list1.Next = mergeTwoLists(list1.Next, list2)
    return list1
  } else {
    list2.Next = mergeTwoLists(list1, list2.Next)
    return list2
  }
}
