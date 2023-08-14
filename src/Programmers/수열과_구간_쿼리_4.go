func solution(arr []int, queries [][]int) []int {
    for _, query := range queries {
        k := query[2]

        for i := query[0]; i <= query[1]; i++ {
            if i % k == 0 {
                arr[i]++
            }
        }
    }

    return arr
}
