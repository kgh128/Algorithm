import "strings"

func solution(str1 string, str2 string) int {
    if strings.Contains(str2, str1) {
        return 1
    }
    return 0
}
