func solution(num_list []int) int {
    var result int

    if len(num_list) >= 11 {
        result = 0

        for _, num := range num_list {
            result += num
        }
    } else {
        result = 1

        for _, num := range num_list {
            result *= num
        }
    }

    return result
}
