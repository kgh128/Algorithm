/*
1. 접두사가 앞에서 n글자짜리의 문자열임을 이용하기
- 접미사: 특정 인덱스까지의 문자열
=> 따라서 접두사는 앞에서 접두사 길이만큼의 문자열이다.
ex) "HelloWorld"에서 "Hello"는 앞에서 5글자 길이의 문자열
=> 앞에서부터 n글자짜리의 문자열은 str[:n]으로 확인한다.
*/

func solution(my_string string, is_prefix string) int {
    if len(is_prefix) > len(my_string) {
        return 0
    }
    if is_prefix == my_string[:len(is_prefix)] {
        return 1
    }
    return 0
}


/*
2. strings의 HasPrefix() 함수 이용하기
- 첫번째 인자: 본래 문자열
- 두번째 인자: 접두사
*/

import "strings"

func solution(my_string string, is_prefix string) int {
    if strings.HasPrefix(my_string, is_prefix) {
        return 1
    }
    return 0
}
