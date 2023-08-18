func solution(board [][]int, k int) int {
    sum := 0
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if i + j <= k {
                sum += board[i][j]
            }
        }
    }
    
    return sum
}
