func solution(x int) []int {
    seq := []int{x}  // int 슬라이스를 만들고, 초기값을 {x}로 함.

    for x != 1 {
        if x%2 == 0 {
            x = x / 2
        } else {
            x = 3 * x + 1
        }

        seq = append(seq, x)
    }

    return seq
}
