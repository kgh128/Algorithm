func solution(n int, k int) []int {
    arr := []int{}

    for i := k; i <= n; i += k {
        arr = append(arr, i)
    }

    return arr
}
