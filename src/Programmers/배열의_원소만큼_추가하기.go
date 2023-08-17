func solution(arr []int) []int {
    X := []int{}

    for _, num := range arr {
        for i := 0; i < num; i++ {
            X = append(X, num)
        }
    }

    return X
}
