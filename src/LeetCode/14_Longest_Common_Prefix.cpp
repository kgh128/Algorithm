// 1. 거장 첫번째 문자열을 기준으로 모든 문자열과 비교하기
// 시간복잡도: O(nm) - n: 문자열 개수, m: 문자열 길이
// (substr도 고려하면 더 커지겠지만, substr은 충분히 push_back을 이용하는 방식으로 변경 가능)
// 공간복잡도: O(m) - m: 가장 짧은 문자열 길이
// 여기에서는 prefix를 구할 때 공통 문자열이 끝나는 인덱스를 구하고, 나중에 prefix를 0부터 해당 인덱스까지 자르는 방식을 사용
#include <string>

class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);

        std::string prefix = strs[0];

        for (int i = 1; i < strs.size(); i++) {
            int j = 0;

            while (j < prefix.size() && j < strs[i].size()) {
                if (prefix[j] != strs[i][j]) {
                    break;
                }
                j++;
            }
            prefix = prefix.substr(0, j);

            if (prefix.size() == 0) {
                return prefix;
            }
        }

        return prefix;
    }
};


// 2. 문자열을 정렬한 후 가장 첫번째 문자열과 가장 마지막 문자열만 비교하기
// 시간복잡도: O(mnlogn) - n: 문자열 개수, m: 첫번째 문자열과 마지막 문자열 중 더 짧은 길이
// 공간복잡도: O(m) - m: 첫번째 문자열과 마지막 문자열 중 더 짧은 길이
// 문자열 벡터를 정렬하면 첫번째 문자열과 마지막 문자열이 가장 짧은 공통 접두사를 가지고 있다.
// 자세한 설명은 14_Longest_Common_Prefix.go 참고
// 여기에서는 루프를 돌면서 같은 문자를 prefix에 push_back 하는 방식을 사용.
#include <string>
#include <algorithm>

class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);

        std::string prefix = "";

        sort(strs.begin(), strs.end());

        for (int i = 0; i < strs.front().size() && i < strs.back().size(); i++) {
            if (strs.front()[i] != strs.back()[i]) {
                break;
            }
            prefix.push_back(strs.front()[i]);
        }

        return prefix;
    }
};
