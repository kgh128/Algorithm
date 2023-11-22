/*
1. 위 -> 아래, 왼쪽 -> 오른쪽 순서로 순회 (첫번째 풀이)
- 시간복잡도: O(n*m) => n: 행의 개수, m: 가장 긴 행의 열의 개수
- 공간복잡도: O(n)
- LeetCode에서 실행 시간이 4793ms이다. 굉장히 길다.
- 메모리는 적게 사용한다.
- 먼저 수행하는 위 -> 아래는 모든 행을 다 순회하면 된다.
- 그 다음 수행하는 왼쪽 -> 오른쪽은 가장 긴 행의 마지막 열까지 순회하면 된다.
*/
func findDiagonalOrder(nums [][]int) []int {
    traversal := []int{}
    longestRow := 0

    for row := 0; row < len(nums); row++ {
        if len(nums[row]) > len(nums[longestRow]) {
            longestRow = row
        }

        curRow := row
        curCol := 0

        for curRow >= 0 {
            if curCol < len(nums[curRow]) {
                traversal = append(traversal, nums[curRow][curCol])
            }

            curRow -= 1
            curCol += 1
        }
    }

    for col := 1; col < len(nums[longestRow]); col++ {
        curRow := len(nums) - 1
        curCol := col

        for curRow >= 0 {
            if curCol < len(nums[curRow]) {
                traversal = append(traversal, nums[curRow][curCol])
            }

            curRow -= 1
            curCol += 1
        }
    }

    return traversal
}
