/*
1. 접미사가 뒤에서 n글자짜리의 문자열임을 이용하기
- 접미사: 특정 인덱스에서부터 시작하는 문자열
=> 따라서 접미사는 뒤에서 접미사 길이만큼의 문자열이다.
ex) "HelloWorld"에서 "World"는 뒤에서 5글자 길이의 문자열
=> 결국 
*/

func solution(my_string string, is_suffix string) int {
    if len(is_suffix) > len(my_string) {
        return 0
    }
    
    if is_suffix == my_string[len(my_string)-len(is_suffix):] {
        return 1
    }
    
    return 0
}
