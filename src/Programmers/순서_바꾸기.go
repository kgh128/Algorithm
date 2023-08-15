func solution(num_list []int, n int) []int {
    return append(num_list[n:], num_list[:n]...)
}
