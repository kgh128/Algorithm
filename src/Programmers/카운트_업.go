func solution(start int, end int) []int {
    result := make([]int, end - start + 1)

    for idx, _ := range result {
        result[idx] = idx + start
    }

    return result
}
