// 1. x 전체를 뒤집어서 reverseX 만들기
// 시간복잡도: O(n), 공간복잡도: O(1)
class Solution {
public:
    bool isPalindrome(int x) {
        ios::sync_with_stdio(0);
        cin.tie(0);
        cout.tie(0);
        
        if (x < 0) {
            return false;
        }

        int copyX = x;
        unsigned int reverseX = 0;

        while (copyX != 0) {
            int rest = copyX % 10;
            copyX /= 10;
            reverseX = reverseX * 10 + rest;
        }

        return x == reverseX;
    }
};


// 2. x를 절반만 뒤집어서 reverseX 만들기
// 시간복잡도: O(n/2), 공간복잡도: O(1)
// +) 애초에 대칭이 안되는 경우 제외: 음수, 일의 자릿수가 0 (0으로 시작하는 숫자는 없음.)
// 자세한 설명은 9_Palindrome_Number.go 파일에 존재
class Solution {
public:
    bool isPalindrome(int x) {
        ios::sync_with_stdio(false);
        cin.tie(NULL);
        
        if (x < 0 || (x % 10 == 0 && x != 0)) {
            return false;
        }

        int reverseX = 0;

        while (x > reverseX) {  // x <= reverseX가 되면 x의 절반을 reverseX로 옮긴 것을 의미
            reverseX = reverseX * 10 + x % 10;
            x /= 10;
        }

        // x가 짝수 개의 숫자로 구성되어 있으면 x와 reverseX는 동일한 절반을 가지고 있으므로 바로 비교 가능
        // x가 홀수 개의 숫자로 구성되어 있으면 reverseX가 가운데 숫자를 포함하여 한 개를 더 가지고 있음.
        // 따라서 어차피 짝이 없는 x의 가운데 숫자(reverseX에서는 일의 자릿수)는 제외하고 x와 reverseX 비교.
        return x == reverseX || x == reverseX / 10;
    }
};
