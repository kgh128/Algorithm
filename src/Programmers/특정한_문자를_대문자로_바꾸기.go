/*
<영소문자 1글자로 이루어진 문자열 alp를 대문자로 바꾸기>
=> 대문자로 바꾼 것도 결과가 문자열이어야 한다.
1. strings.ToUpper(alp)
2. string(alp[0] - 32)
   -> alp[0]을 이용하여 1글자 문자열을 byte 타입으로 접근한다.
   -> 그냥 byte(alp)를 하면 컴파일 에러가 난다.
*/

import "strings"

func solution(my_string string, alp string) string {
    return strings.ReplaceAll(my_string, alp, strings.ToUpper(alp))
}
