import "bytes"

func solution(code string) string {
    mode := 0
    var ret bytes.Buffer
    
    for idx, char := range code {
        if char == '1' {
            mode = 1 - mode        // mode: 0 -> 1, 1 -> 0
        } else if idx%2 == mode {  // mode가 0이면 idx가 짝수, mode가 1이면 idx가 홀수
            ret.WriteRune(char)
        }
    }
    
    if ret.String() == "" {
        return "EMPTY"
    } else {
        return ret.String()
    }
}
