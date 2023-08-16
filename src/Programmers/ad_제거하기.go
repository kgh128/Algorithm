import "strings"

func solution(strArr []string) []string {
    result := []string{}

    for _, str := range strArr {
        if !strings.Contains(str, "ad") {
            result = append(result, str)
        }
    }

    return result
}
