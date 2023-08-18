import "strings"

func solution(my_string string, target string) int {
    if strings.Contains(my_string, target) {
        return 1
    }
    return 0
}
