func solution(a int, d int, included []bool) int {
    sum := 0

    for idx, isTrue := range included {
        if isTrue {
            sum += (a + d * idx)  // 등차수열의 항
        }
    }

    return sum
}
