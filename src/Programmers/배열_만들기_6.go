func solution(arr []int) []int {
    stk := []int{}
    
    for i := 0; i < len(arr); i++ {
        if len(stk) > 0 && stk[len(stk)-1] == arr[i] {
            stk = stk[:len(stk)-1]
        } else {
            stk = append(stk, arr[i])
        }
    }
    
    if len(stk) == 0 {
        stk = append(stk, -1)
    }
    
    return stk
}
