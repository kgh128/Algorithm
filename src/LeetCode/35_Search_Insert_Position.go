// 1. 이진탐색으로 풀기 -> 시간: 4ms
// 이진탐색을 재귀 함수로 구현하여 오히려 시간이 더 걸리는 것일 수도 있다.
func searchInsert(nums []int, target int) int {
    midIdx := len(nums) / 2

    if nums[midIdx] == target {
        return midIdx
    } else if nums[midIdx] > target {
        if midIdx == 0 {
            return midIdx
        }
        return searchInsert(nums[:midIdx], target)
    } else {
        if midIdx == len(nums) - 1 {
            return midIdx + 1
        }
        return midIdx + 1 + searchInsert(nums[midIdx+1:], target)
    }
}

// 2. 선형탐색으로 풀기 -> 시간: 0ms
func searchInsert(nums []int, target int) int {
    for i, num := range(nums) {
        if num >= target {
            return i
        }
    }
    return len(nums)
}
