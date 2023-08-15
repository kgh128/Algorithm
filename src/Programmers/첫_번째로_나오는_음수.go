func solution(num_list []int) int {
    for idx, num := range num_list {
        if (num < 0) {
            return idx
        }
    }
    return -1
}
