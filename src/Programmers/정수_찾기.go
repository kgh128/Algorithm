func solution(num_list []int, n int) int {
    for _, num := range num_list {
        if num == n {
            return 1
        }
    }
    return 0
}
