func solution(arr []int) []int {
    stk := []int{}

    for i := 0; i < len(arr); {
        end := len(stk) - 1

        if len(stk) == 0 || stk[end] < arr[i] {
            stk = append(stk, arr[i])
            i++
        } else {
            stk = stk[:end]
        }
    }

    return stk
}
