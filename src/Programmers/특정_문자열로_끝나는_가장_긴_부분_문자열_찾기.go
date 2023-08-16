/*
<strings.LastIndex() 이용>
- 첫번째 인자: 원본 문자열
- 두번째 인자: 찾고 싶은 특정 문자열
- 원본 문자열에서 가장 마지막에 존재하는 특정 문자열의 시작 인덱스를 반환한다.
*/

import "strings"

func solution(myString string, pat string) string {
    end := strings.LastIndex(myString, pat) + len(pat)
    return myString[:end]
}
