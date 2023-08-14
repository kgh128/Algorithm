/*
1. 전체 문자열을 복사하고, s부터 e까지만 뒤집기
-> swap 사용
-> 구간 길이의 절반을 기준으로 문자를 서로 swap
ex) "12ab"는 1 2 | a b로 절반으로 나눌 수 있음.
    1 <-> b, 2 <-> a 이렇게 swap
*/

func solution(my_string string, s int, e int) string {
    bytes := []byte(my_string)
    half := (e - s + 1) / 2

    for i := 0; i < half; i++ {  // (구간 길이)/2만큼 for문을 돎.
        bytes[s+i], bytes[e-i] = bytes[e-i], bytes[s+i]
    }

    return string(bytes)
}


/*
2. s부터 e까지만 뒤집어서 중간 문자열을 만들고, 앞뒤에 본래 문자열 붙여주기
-> 처음부터 전체 복사열을 복사할 필요X
-> for문만 거꾸로 돌면서 문자열을 뒤집어서 복사
*/
import "bytes"

func solution(my_string string, s int, e int) string {
    var buf bytes.Buffer
    
    buf.WriteString(my_string[:s])  // 구간 앞 문자열 추가
    
    for i := e; i >= s; i-- {  // (구간 길이)만큼 for문을 돎.
        buf.WriteByte(my_string[i])  // 뒤집은 구간 추가
    }
    
    buf.WriteString(my_string[e+1:])  // 구간 뒷 문자열 추가
    
    return buf.String()
}
