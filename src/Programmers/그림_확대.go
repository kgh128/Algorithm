func solution(picture []string, k int) []string {
    expandPic := []string{}

    for _, row := range picture {
        expandRow := []rune{}

        for _, pixel := range row {
            for i := 0; i < k; i++ {
                expandRow = append(expandRow, pixel)
            }
        }

        for i := 0; i < k; i++ {
            expandPic = append(expandPic, string(expandRow))
        }
    }

    return expandPic
}
