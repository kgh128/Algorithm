func solution(arr []int, flags []bool) []int {
    x := []int{}

    for i, flag := range flags {
        if flag {
            for j := 0; j < arr[i] * 2; j++ {
                x = append(x, arr[i])
            }
        } else {
            x = x[:len(x)-arr[i]]
        }
    }

    return x
}
