func solution(arr []int, k int) []int {
    result := make([]int, 0, k)
    flag := [100001]bool{}      // 전체 배열이 false로 초기화됨.
                                // flag[n]이 true이면 n은 이미 result에 추가된 것.
                                // flag := make(map[int]bool) 
                                // 이렇게 key가 int이고, value가 bool인 map의 슬라이스 형태로도 만들 수 있음.

    for i := 0; len(result) < k; i++ {
        if i >= len(arr) {
            result = append(result, -1)
        } else if !flag[arr[i]] {
            result = append(result, arr[i])
            flag[arr[i]] = true
        }
    }

    return result
}
