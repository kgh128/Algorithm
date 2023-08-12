/*
<Atoi(), Itoa() 이용 / '+' 이용하여 문자열 이어붙이기>
- Atoi(): string(아스키코드) -> int
- Itoa(): int -> string(아스키코드)
*/

import "strconv"

func solution(a int, b int) int {
    ab, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
    ba, _ := strconv.Atoi(strconv.Itoa(b) + strconv.Itoa(a))
    
    if ab > ba {
        return ab
    } else {
        return ba
    }
}
