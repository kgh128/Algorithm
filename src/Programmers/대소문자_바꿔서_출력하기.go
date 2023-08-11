/*
<아스키코드 이용하기> - 여기서는 10진수로 이용
'A' - 'Z': 65 - 90
'a' - 'z': 97 - 122

따라서 대소문자의 아스키코드는 32의 차이가 남.
대문자 + 32 = 소문자

<string을 range를 이용하여 원소 순회>
원소의 타입: rune -> 유니코드에서 문자의 코드 포인트를 저장함.
(알파벳이면 아스키코드 값이 들어있음.) (int의 alias -> 상호호환)
ex) 'A'이면 65가 들어있음.
    따라서 rune 타입 변수에 'A' + 32를 넣고 해당 변수를 출력하면 아스키코드가 97인 'a'가 출력됨.
*/

package main

import "fmt"

func main() {
    var str string
    fmt.Scan(&str)
    
    retStr := make([]rune, len(str))

    for i, char := range str {
        // 'Z'의 아스키코드가 90이므로, 아스키코드가 91보다 작으면 대문자
        if char < 91 {
            retStr[i] = char + 32  // 대문자 -> 소문자
        } else {
            retStr[i] = char - 32  // 소문자 -> 대문자
        }
    }
    
    fmt.Println(string(retStr))
}
