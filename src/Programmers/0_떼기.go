/*
1. 처음으로 0이 아닌 값이 나오면 해당 원소의 인덱스부터 끝까지 문자열 자르기
*/
func solution(n_str string) string {
    start := 0

    for _, n := range n_str {
        if n != '0' {
            break
        }
        start++
    }

    return n_str[start:]
}


/*
2. strings.TrimLeft() 이용
- 문자열 앞에 오는 특정 문자열을 제거해주는 함수
- 첫번째 인자: 원본 문자열
- 두번째 인자: 제거할 특정 문자열
*/
import "strings"

func solution(n_str string) string {
    return strings.TrimLeft(n_str, "0")
}
