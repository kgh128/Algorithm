func removeDuplicates(nums []int) int {
    // 1. 애초에 배열 원소의 개수가 0이나 1이면 중복된 원소가 있을 리가 없으니 바로 리턴.
    if len(nums) < 2 {
        return len(nums)
    }

    // 2. 중복 원소가 아닌 원소들 중에서 가장 마지막 인덱스
    latestIdx := 0

    // 3. 반복문 돌면서 중복 원소 지우기
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[latestIdx] {
            latestIdx++
            nums[latestIdx] = nums[i]
        }
    }

    // 4. 인덱스에 1을 더하면 길이
    return latestIdx + 1
}
