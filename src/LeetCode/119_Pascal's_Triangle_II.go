func getRow(rowIndex int) []int {
    pascal := make([]int, rowIndex+1)

    for i := 0; i <= rowIndex; i++ {
        for j := i; j >= 0; j-- {
            if j == 0 || j == i {
                pascal[j] = 1
            } else {
                pascal[j] += pascal[j-1]
            }
        }
    }

    return pascal
}
