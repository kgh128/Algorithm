/*
<num - rev(num)을 해시 테이블에 저장하기>
- 시간복잡도: O(n), 공간복잡도: O(n)
- 원래 조건: nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
- 이항해서 정리: nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
  => 루프를 돌면서 각 num에 따라 num - rev(num)을 구하면 된다.
  => 기존 조건 그대로 루프에서 비교하려면 i와 j에 대해서 이중 for문을 돌아야 했다. -> 시간복잡도: O(n^2)
  => 바뀐 조건대로 하면 i에 대해서만 for문을 돌면서 수식을 계산하면 된다.
  => 계산한 값은 해시 테이블에 저장하여 비교할 때 사용하면 된다. -> 다른 num과 비교하기 위해 루프를 돌지 않아도 된다.
- golang은 각 타입마다 zero value가 존재한다.
  => map[int]int에서 gapCount[gap]의 값을 할당하지 않고 바로 사용하면 gapCount[gap]의 값은 int형의 zero value인 0이 된다.
  => 꼭 if문으로 map에 해당 키가 존재하는지 검사하지 않아도 사용 가능하다.
  => 그 gapCount[gap]에 += 연산자나 ++ 연산자 모두 사용 가능하다.
- 문제에 "Since that number can be too large, return it modulo 10^9 + 7"와 같은 조건이 존재하므로 count 값은 항상 1000000007의 나머지 값으로 유지한다.
*/
func countNicePairs(nums []int) int {
    mod := 1000000007
    count := 0
    gapCount := make(map[int]int)

    for _, num := range nums {
        gap := num - rev(num)
        count = (count + gapCount[gap]) % mod
        gapCount[gap] += 1
    }

    return count
}

func rev(num int) int {
    revNum := 0

    for num > 0 {
        revNum = revNum * 10 + num % 10
        num /= 10
    }

    return revNum
}
