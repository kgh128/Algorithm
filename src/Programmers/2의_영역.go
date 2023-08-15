func solution(arr []int) []int {
    var start, end = -1, -1

    for idx, num := range arr {
        if num == 2 {
            if start == -1 {
                start = idx
            }
            end = idx
        }
    }

    if start == -1 {
        return []int{-1}
    } else {
        return arr[start:end+1]
    }
}
