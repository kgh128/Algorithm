func solution(arr [][]int) int {
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr); j++ {
            if arr[i][j] != arr[j][i] {
                return 0
            }
        }
    }
    return 1
}
