func solution(arr []int, queries [][]int) []int {
    for _, query := range queries {
        for i := query[0]; i <= query[1]; i++ {
            arr[i]++
        }
    }
    return arr
}
