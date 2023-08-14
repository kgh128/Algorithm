/*
1. 첫번째 행의 c번째 열부터 시작함. (i는 0부터 시작X)
=> 단, 배열의 첫번째 인덱스는 0이므로 실제로는 i가 c-1부터 시작
2. 그 후 인덱스를 m씩 늘려주면 i는 그 다음 행의 c번째 열을 가리킴.

ex) 0 1 2 3
    4 5 6 7
    8 9 10 11
    이렇게 있고, m = 4, c = 3이면
    애초에 2부터 시작하면 m씩 늘어나니까 그 다음에는 바로 6, 그 다음에는 바로 10으로 접근 가능.
*/

func solution(my_string string, m int, c int) string {
    column := []byte{}
    
    for i := c-1; i < len(my_string); i += m {
        column = append(column, my_string[i])
    }
    
    return string(column)
}
