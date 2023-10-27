/*
<strings.ToLoWer() 이용해서 문자열의 모든 대문자를 소문자로 바꾸기>
- 문제에서 대문자와 소문자를 구분하지 않는다.
=> 대문자, 소문자 상관없이 같은 문자인지 검사해야 한다.
=> 그러나 실제 Go에서는 같은 문자인지 검사할 때, 대문자와 소문자를 구분한다.
=> 따라서 문자열의 모든 대문자를 소문자로 바꾸어 문자를 검사할 때는 소문자끼리 비교하도록 한다.

- strings.ToLower() 이용
*/


// 1. strings.Contain() 이용하여 특정 패턴이 문자열에 존재하는지 확인
// - 첫번째 인자: 문자열
// - 두번째 인자: 패턴

import "strings"

func solution(myString string, pat string) int {
    if strings.Contains(strings.ToLower(myString), strings.ToLower(pat)) {
        return 1
    }
    return 0
}


// 2. 반복문 돌면서 패턴의 길이만큼 부분 문자열을 잘라서 검사하기
// - myString[i:i+len(pat)]가 pat와 같은지 확인한다.
// - 그러면 마지막 부분 문자열은 인덱스가 len(myString) - len(pat)에서 시작한다.

import "strings"

func solution(myString string, pat string) int {
    myString = strings.ToLower(myString)
    pat = strings.ToLower(pat)
    
    for i := 0; i <= len(myString) - len(pat); i++ {
        if (myString[i:i+len(pat)] == pat) {
            return 1
        }
    }
    return 0
}
