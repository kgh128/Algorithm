/*
<bytes.Buffer 사용>
- WriteString() 함수를 이용하여 버퍼에 문자열을 이어붙일 수 있다.
- String() 함수를 이용하여 bytes -> string으로 변환할 수 있다.
=> 이어붙일 문자열이 많을수록 '+' 연산자보다 좋은 성능을 낸다.

이것 말고도 string slice를 만들고, Join() 함수를 통해 문자열을 이용하여 이어붙이는 방법도 존재한다.
*/

import "bytes"

func solution(my_string string, overwrite_string string, s int) string {
    var buf bytes.Buffer
    
    buf.WriteString(my_string[:s])
    buf.WriteString(overwrite_string)
    buf.WriteString(my_string[s + len(overwrite_string):])
  
    // s <= len(my_string) - len(overwrite_string)이므로
    // s + len(overwrite_string) <= len(my_string)
    // 문자열을 [m:n]을 이용하여 자를 때, m과 n에 올 수 있는 최댓값은 해당 문자열의 길이
    // ex) len(my_string)
    // 왜냐하면 문자열의 끝까지 자를 때는 n 자리에 (끝 인덱스) + 1이 와야하는데, 그게 문자열의 길이
    // m과 n이 동일하면 결과 문자열은 빈 문자열(시작과 끝이 같으니까 자를 게 없음.) -> 즉, m의 최댓값도 n의 최댓값과 동일
    // 따라서 이 문제에서는 s + len(overwrite_string)이 인덱스 범위를 벗어나는지 체크할 필요X
    
    return buf.String()
}
