func solution(strArr []string) int {
    group := make(map[int]int)
    maxLength := 0
    
    for _, str := range strArr {
        group[len(str)] += 1
    }

    // map에 있는 값 중 가장 큰 값 찾기
    // 이걸 위해서 굳이 map을 정렬하려고 할 필요X
    for _, length := range group {
        if (length > maxLength) {
            maxLength = length
        }
    }
    
    return maxLength
}
