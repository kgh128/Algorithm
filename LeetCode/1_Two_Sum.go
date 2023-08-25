// 1. 완전탐색 => Time: O(n^2), Space: O(1)
func twoSum(nums []int, target int) []int {
    for i, num := range nums {
        for j := i + 1; j < len(nums); j++ {
            if num + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}

// 2. Map 이용 => Time: O(n), Space: O(n)
// map에 key: num, value: index로 저장
// nums를 반복문을 돌면서 num을 하나씩 순회하는데, num과 더해서 target이 되려면 필요한 수는 target-num
// map에 target-num이 존재하면 사용할 수 있는 수이므로 num과 target-num의 index를 리턴한다.
// 존재하지 않으면 현재 순회한 num과 그 index를 map에 저장하여 나중에 검색할 수 있도록 한다.

// if문에 간단한 문장(optional statement)를 함깨 쓸 수 있다.
// map에 해당 key가 존재하는지는 map[key]의 두번째 리턴값을 가지고 확인할 수 있다.
// map[key]의 첫번째 리턴값: value, 두번째 리턴값: 해당 key가 존재하는지를 의미하는 boolean 값
func twoSum(nums []int, target int) []int {
    indices := make(map[int]int)

    for i, num := range nums {
        if idx, isPresent := indices[target-num]; isPresent {
            return []int{i, idx}
        }
        indices[num] = i
    }

    return []int{}
}
