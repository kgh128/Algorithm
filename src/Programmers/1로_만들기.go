func solution(num_list []int) int {
    count := 0

    for _, num := range num_list {
        for num != 1 {
            if num % 2 == 0 {
                num /= 2
            } else {
                num = (num - 1) / 2
            }

            count++
        }
    }

    return count
}
