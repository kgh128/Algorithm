import "strings"

func solution(myStr string) []string {
    strArr := strings.FieldsFunc(myStr, func(r rune) bool {
        return r == 'a' || r == 'b' || r == 'c'
    })

    if len(strArr) == 0 {
        strArr = append(strArr, "EMPTY")
    }

    return strArr
}
