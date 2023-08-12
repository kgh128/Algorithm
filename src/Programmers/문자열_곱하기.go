import "strings"

// 1. strings.Builder 이용해서 문자열 반복해서 이어붙이기
func solution(my_string string, k int) string {
    var ret_string strings.Builder

    for i := 0; i < k; i++ {
        ret_string.WriteString(my_string)
    }

    return ret_string.String()
}

// 2. strings.Repeat 이용해서 문자열 반복해서 이어붙이기
func solution(my_string string, k int) string {
    return strings.Repeat(my_string, k)
    // my_string을 k번 반복한다는 의미
}
