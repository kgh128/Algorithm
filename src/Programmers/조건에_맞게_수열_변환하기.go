func solution(arr []int) []int {
    for idx, num := range arr {
        if num >= 50 && num % 2 == 0 {
            arr[idx] /= 2
        } else if num < 50 && num % 2 == 1 {
            arr[idx] *= 2
        }
    }
    return arr
}
