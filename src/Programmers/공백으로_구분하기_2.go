/*
<strings.Fields()>
- 공백이 연속해서 한 개 이상 존재해도 구분 가능
ex) "a b c d"와 "a    b   c d" 모두 결과는 ["a", "b", "c", "d"]
*/

import "strings"

func solution(my_string string) []string {
    return strings.Fields(my_string)
}
