import "strings"

func solution(str_list []string, ex string) string {
    tail := []string{}

    for _, str := range str_list {
        if !strings.Contains(str, ex) {
            tail = append(tail, str)
        }
    }

    return strings.Join(tail, "")
}
