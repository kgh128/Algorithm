func solution(arr []int, queries [][]int) []int {
    results := make([]int, len(queries))

    for idx, query := range queries {
        results[idx] = 1000001

        for i := query[0]; i <= query[1]; i++ {
            if (arr[i] > query[2] && arr[i] < results[idx]) {
                results[idx] = arr[i]
            }
        }

        if results[idx] == 1000001 {
            results[idx] = -1
        }
    }

    return results
}
