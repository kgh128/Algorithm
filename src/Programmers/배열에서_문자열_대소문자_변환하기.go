import "strings"

func solution(strArr []string) []string {
    for idx, _ := range strArr {
        if (idx % 2 == 1) {
            strArr[idx] = strings.ToUpper(strArr[idx])
        } else {
            strArr[idx] = strings.ToLower(strArr[idx])
        }
    }
    return strArr
}
