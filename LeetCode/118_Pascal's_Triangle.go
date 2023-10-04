func generate(numRows int) [][]int {
    pascal := make([][]int, numRows)

    for i := 0; i < numRows; i++ {
        pascal[i] = make([]int, i + 1)

        for j := 0; j < i + 1; j++ {
            if j == 0 || j == i {
                pascal[i][j] = 1
            } else {
                pascal[i][j] = pascal[i-1][j-1] + pascal[i-1][j]
            }
        }
    }

    return pascal
}
