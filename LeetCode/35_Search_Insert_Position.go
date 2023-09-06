// 1. 이진탐색으로 풀기 -> 시간: 4ms
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

// 2. 선형탐색으로 풀기 
