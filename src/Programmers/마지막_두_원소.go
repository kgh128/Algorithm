func solution(nums []int) []int {
    end := len(nums) - 1

    if nums[end] > nums[end-1] {
        return append(nums, nums[end] - nums[end-1])
    } else {
        return append(nums, nums[end] * 2)
    }
}
