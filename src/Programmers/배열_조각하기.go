func solution(arr []int, queries []int) []int {
    for idx, query := range queries {
        if idx % 2 == 0 {
            arr = arr[:query+1]
        } else {
            arr = arr[query:]
        }
    }
    return arr
}
