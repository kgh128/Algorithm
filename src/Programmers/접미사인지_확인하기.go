/*
1. 접미사가 뒤에서 n글자짜리의 문자열임을 이용하기
- 접미사: 특정 인덱스에서부터 시작하는 문자열
=> 따라서 접미사는 뒤에서 접미사 길이만큼의 문자열이다.
ex) "HelloWorld"에서 "World"는 뒤에서 5글자 길이의 문자열
=> 뒤에서부터 n글자짜리의 문자열은 str[len(str)-n:]으로 확인한다.
*/

func solution(my_string string, is_suffix string) int {
    // 접미사의 길이가 본래 문자열보다 긴 경우는 당연히 본래 문자열의 접미사가 될 수 없다.
    // 이를 먼저 확인해줘야만, 아래의 my_string[len(my_string)-len(is_suffix):] 코드에서
    // 인덱스 범위를 벗어났다는 오류가 나지 않는다.
    // 이 경우에는 len(my_string) - len(is_suffix) < 0이기 때문이다.
    if len(is_suffix) > len(my_string) {
        return 0
    }
    
    if is_suffix == my_string[len(my_string)-len(is_suffix):] {
        return 1
    }
    
    return 0
}


/*
2. strings의 HasSuffix() 함수 이용하기
- 첫번째 인자: 본래 문자열
- 두번째 인자: 접미사
*/

import "strings"

func solution(my_string string, is_suffix string) int {
    if strings.HasSuffix(my_string, is_suffix) {
        return 1
    }
    return 0
}
