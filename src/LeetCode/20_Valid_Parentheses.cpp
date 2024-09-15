// 1. Stack을 이용해서 풀기
// 시간복잡도: O(n), 공간복잡도: O(n)
// 여는 괄호는 stack에 집어넣고, 닫는 괄호인 경우에는 stack에서 pop하여 짝이 맞는 여는 괄호인지 확인
// +) 여는 괄호인지 확인할 때 문자 비교를 일일이 하지 않고 pair map의 key에 해당 값이 존재하는지 확인해도 된다.
#include <stack>
#include <unordered_map>

class Solution {
public:
    bool isValid(string s) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);

        stack<char> stack;
        unordered_map<char, char> pair = {
            {'(', ')'},
            {'{', '}'},
            {'[', ']'}
        };

        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '(' || s[i] == '{' || s[i] == '[') {
                stack.push(s[i]);
            } else {
                if (stack.empty() || pair[stack.top()] != s[i]) {
                    return false;
                }
                stack.pop();
            }
        }

        return stack.empty();
    }
};
