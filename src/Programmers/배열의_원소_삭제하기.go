func isNotDeleted(num int, deleteList []int) bool {
    for _, deleteNum := range deleteList {
        if num == deleteNum {
            return false
        }
    }
    return true
}

func solution(arr []int, deleteList []int) []int {
    rest := []int{}

    for _, num := range arr {
        if isNotDeleted(num, deleteList) {
            rest = append(rest, num)
        }
    }

    return rest
}
