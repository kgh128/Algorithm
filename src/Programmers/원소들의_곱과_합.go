func solution(num_list []int) int {
    multiply := 1
    sum := 0

    for _, num := range num_list {
        multiply *= num
        sum += num
    }

    if multiply < (sum * sum) {
        return 1
    } else {
        return 0
    }
}
