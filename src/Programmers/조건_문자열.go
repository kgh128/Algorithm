/*
<switch문 사용>
1. 비교값에 식이 들어갈 수 있음.
ex) switch ineq + eq {...}

2. 초기화문도 사용가능 (초기화문 다음에 비교값이 들어가야함.)
ex) switch op := ineq + eq; op {...}
*/

func solution(ineq string, eq string, n int, m int) int {
    isTrue := false
    
    switch ineq + eq {
        case ">=":
            isTrue = (n >= m)
        case "<=":
            isTrue = (n <= m)
        case ">!":
            isTrue = (n > m)
        case "<!":
            isTrue = (n < m)
    }
    
    if isTrue {
        return 1
    } else {
        return 0
    }
}
