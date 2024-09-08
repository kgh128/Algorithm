// 1. 완전탐색 => Time: O(n^2), Space: O(1)
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> result;

        for (int i = 0; i < nums.size(); i++) {
            for (int j = i+1; j < nums.size(); j++) {
                if (nums[i] + nums[j] == target) {
                    result.push_back(i);
                    result.push_back(j);
                    return result;
                }
            }
        }
        return result;
    }
};


// 2. Map 이용 => Time: O(n), Space: O(n)
// map에 key: num, value: index로 저장
// nums를 반복문을 돌면서 num을 하나씩 순회하는데, num과 더해서 target이 되려면 필요한 수는 target-num
// map에 target-num이 존재하면 사용할 수 있는 수이므로 num과 target-num의 index를 리턴한다.
// 존재하지 않으면 현재 순회한 num과 그 index를 map에 저장하여 나중에 검색할 수 있도록 한다.

// 32-34 라인: 한 번에 묶어서 입출력
// {}로 vector를 바로 만들기 가능
// contains() 함수로 key가 map에 존재하는지 확인 가능
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        ios::sync_with_stdio(0);
        cin.tie(0);
        cout.tie(0);

        unordered_map<int, int> index;

        for (int i = 0; i < nums.size(); i++) {
            int num = target - nums[i];

            if (index.contains(num)) {
                return {i, index[num]};
            }
            index[nums[i]] = i;
        }
        return {-1};
    }
};
