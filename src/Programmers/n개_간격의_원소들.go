func solution(num_list []int, n int) []int {
    result := []int{}

    for i := 0; i < len(num_list); i += n {
        result = append(result, num_list[i])
    }

    return result
}
