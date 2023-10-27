func climbStairs(n int) int {
    if n <= 2 {
        return n
    }

    stair := make([]int, n + 1)  // i칸까지 오는 경우의 수를 저장한다. -> dp에 이용

    stair[1] = 1
    stair[2] = 2

    for i := 3; i <= n; i++ {
        stair[i] = stair[i-2] + stair[i-1]
        // 2칸 아래에서 한 번에 올 수 있으니까 stair[i-2]까지 오는 경우의 수를 더하면 된다.
        // 1칸 아래에서 한 번에 올 수 있으니까 stair[i-1]까지 오는 경우의 수를 더하면 된다.
    }

    return stair[n]
}
