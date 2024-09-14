// 1. 문자 비교로 특수 경우 탐지
// 시간복잡도: O(n), 공간복잡도: O(1)
// 문자열을 앞에서부터 순회
// 현재 문자가 특수 경우에서 앞에 오는 문자라면 바로 다음 문자가 특수 경우에서 뒤에 오는 문자인지 확인
// 확인해서 특수 경우이면 해당 값 더하고, 다다음 문자로 넘어가기
// +) 모든 특수 경우에 대해 switch + if문 작성
#include <unordered_map>

class Solution {
public:
    int romanToInt(string s) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);

        int result = 0;

        std::unordered_map<char, int> intMap =
        {
            {'I', 1},
            {'V', 5},
            {'X', 10},
            {'L', 50},
            {'C', 100},
            {'D', 500},
            {'M', 1000}
        };

        for (int i = 0; i < s.size(); i++) {
            if (i < s.size() - 1) {
                switch (s[i]) {
                    case 'I':
                        if (s[i+1] == 'V') {result += 4; i++; continue;}
                        if (s[i+1] == 'X') {result += 9; i++; continue;}
                        break;
                    case 'X':
                        if (s[i+1] == 'L') {result += 40; i++; continue;}
                        if (s[i+1] == 'C') {result += 90; i++; continue;}
                        break;
                    case 'C':
                        if (s[i+1] == 'D') {result += 400; i++; continue;}
                        if (s[i+1] == 'M') {result += 900; i++; continue;}
                        break;
                }
            }
            result += intMap[s[i]];
        }

        return result;
    }
};


// 2. 숫자 비교로 특수 경우 탐지
// 시간복잡도: O(n), 공간복잡도: O(1)
// 문자열을 뒤에서부터 순회
// 일반적인 경우에는 현재 값이 직전 값보다 커야 함. 문자열 앞으로 갈수록 큰 숫자를 표기하기 때문이다.
// 특수 경우에는 현재 값이 직전 값보다 작음.
// 이 경우에는 total에서 현재 값을 뺌. 그러면 total 입장에서는 (직전 값 - 현재 값)을 더한 것임. (이미 이전 루프에서 직전 값을 더했기 때문.)
// +) 특수 경우(AB)의 값: A < B, (B - A)를 의미.
#include <unordered_map>

class Solution {
public:
    int romanToInt(string s) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);

        int curr = 0;
        int prev = 0;
        int total = 0;

        std::unordered_map<char, int> intMap =
        {
            {'I', 1},
            {'V', 5},
            {'X', 10},
            {'L', 50},
            {'C', 100},
            {'D', 500},
            {'M', 1000}
        };

        for (int i = s.size() - 1; i >= 0; i--) {
            curr = intMap[s[i]];

            if (curr < prev) {
                total -= curr;
            } else {
                total += curr;
            }

            prev = curr; 
        }

        return total;
    }
};
