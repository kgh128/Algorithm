// 1. 재귀함수로 풀기
// 시간복잡도: O(m+n) - m: list1의 길이, n: list2의 길이
// 공간복잡도: O(m+n) - m: list1의 길이, n: list2의 길이
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);

        if (list1 == NULL) {
            return list2;
        }
        if (list2 == NULL) {
            return list1;
        }

        if (list1->val < list2->val) {
            list1->next = mergeTwoLists(list1->next, list2);
            return list1;
        } else {
            list2->next = mergeTwoLists(list1, list2->next);
            return list2;
        }
    }
};


// 2. 반복문으로 풀기
// 시간복잡도: O(m+n) - m: list1의 길이, n: list2의 길이
// 공간복잡도: O(m+n) - m: list1의 길이, n: list2의 길이
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);

        ListNode* result = new ListNode();
        ListNode* curr = result;

        while (list1 != NULL && list2 != NULL) {
            if (list1->val < list2->val) {
                curr->next = list1;
                list1 = list1->next;
            } else {
                curr->next = list2;
                list2 = list2->next;
            }
            curr = curr->next;
        }

        if (list1 != NULL) {
            curr->next = list1;
        } else if (list2 != NULL) {
            curr->next = list2;
        }

        return result->next;
    }
};
