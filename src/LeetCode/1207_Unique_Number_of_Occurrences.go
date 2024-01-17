/*
<map을 이용하여 빈도도 고유한지 체크>
- 시간복잡도: O(n), 공간복잡도: O(n)
- 처음에는 map을 이용하여 arr 속에서 특정 숫자가 몇 번 나오는지(빈도) 저장
- 두번째는 map의 key를 빈도로 설정하여 같은 빈도가 존재하는지 확인
  => 만약 첫번째 map에서 같은 빈도를 가지고 있는 숫자들이 있었다면
  => 두번쨰 map을 순회할 때 한 숫자는 두번째 map에 저장되어 있을 것이다.
  => 빈도를 key로 넣었을 때 값이 존재한다면 이는 해당 빈도를 가지는 숫자가 이미 존재한다는 의미이므로 고유하지 않다는 의미
*/

func uniqueOccurrences(arr []int) bool {
    countMap := make(map[int]int)
    reverseMap := make(map[int]int)

    for _, val := range arr {
        countMap[val]++
    }

    for val, count := range countMap {
        if _, exists := reverseMap[count]; exists {
            return false
        }
        reverseMap[count] = val
    }

    return true
}
