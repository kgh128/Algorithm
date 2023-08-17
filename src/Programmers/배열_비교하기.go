func solution(arr1 []int, arr2 []int) int {
    if len(arr1) > len(arr2) {
        return 1
    } else if len(arr1) < len(arr2) {
        return -1
    } else {
        sum1 := 0
        sum2 := 0

        for i, _ := range arr1 {
            sum1 += arr1[i]
            sum2 += arr2[i]
        }

        if sum1 > sum2 {
            return 1
        } else if sum1 < sum2 {
            return -1
        } else {
            return 0
        }
    }
}
