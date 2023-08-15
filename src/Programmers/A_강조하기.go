/*
1. 반복문 돌면서 직접 조건에 맞게 바꾸기
- 'A'를 제외한 대문자는 아스키코드의 범위를 이용하여 판별하였다.
+) 빈 byte 슬라이스를 만들고, 반복문 돌면서 조정한 문자를 넣는 방식도 있다.
*/

func solution(myString string) string {
    myStr := []byte(myString)  // 여기서는 byte 슬라이스에 문자열을 그대로 복사
    
    for i, char := range myStr {
        if char == 'a' {
            myStr[i] = 'A'
        } else if char > 'A' && char <= 'Z' {
            myStr[i] += 32
        }
    }
    
    return string(myStr)
}


/*
2. strings의 메소드 이용하기
- strings.ToLower()를 이용하여 문자열의 모든 문자를 소문자로 만들기
=> 어차피 최종 문자열에서 대문자는 "A"밖에 없을 것이므로 일단 모두 소문자로 바꾼다. 
- strings.ReplaceAll()을 이용하여 문자열의 모든 "a"를 "A"로 바꾸기
*/

import "strings"

func solution(myString string) string {
    myString = strings.ToLower(myString)
    return strings.ReplaceAll(myString, "a", "A")
}
