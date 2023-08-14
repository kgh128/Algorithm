func solution(my_string string, queries [][]int) string {
    str := []byte(my_string)

    for _, query := range queries {
        start := query[0]
        end := query[1]

        // 문자열 뒤집기 -> Go에서의 swap이용
        // 중간((end - start + 1)/2)을 기준으로 대칭되는 문자끼리 swap
        // ex) "ab12"이면 a b | 1 2로 나눠지고, a <-> 2, b <-> 1
        // => "21ba"
        for i := 0; i < (end - start + 1)/2; i++ {
            str[start+i], str[end-i] = str[end-i], str[start+i]
        }
    }

    return string(str)
}
