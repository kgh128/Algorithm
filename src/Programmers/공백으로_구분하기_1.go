/*
1. strings.Fields() 이용하기
- 공백을 기준으로 문자열을 잘라서 문자열 슬라이스로 만들어준다.

2. strings.Split() 이용하기
- 특정 문자열을 기준으로 문자열을 잘라서 문자열 슬라이스로 만들어준다.
*/

import "strings"

func solution(my_string string) []string {
    return strings.Fields(my_string)
}
