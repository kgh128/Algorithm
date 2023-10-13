// 1. 소문자 alphanumeric 조건에 맞게 새 배열을 만들고, 그 배열에서 palindrome인지 검사
// 시간: O(n) - 0ms, 공간: O(n) - 3.1MB
func isPalindrome(s string) bool {
    result := []byte{}

    for i := 0; i < len(s); i++ {
        if (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) {
            result = append(result, s[i])
        } else if s[i] >= 65 && s[i] <= 90 {
            result = append(result, s[i] + 32)
        }
    }

    for i := 0; i < len(result) / 2; i++ {
        if result[i] != result[len(result) - 1 - i] {
            return false
        }
    }
    return true
}


// 2. 새 배열을 만들지 않고, 투 포인터로 기존 배열에서 palindrome 검사
