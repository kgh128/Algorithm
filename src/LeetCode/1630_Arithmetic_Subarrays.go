/*
1. 정렬 이용
- 시간복잡도: O(m*nlogn) -> n은 nums 개수, m은 쿼리 개수, logn은 정렬 시간
- 공간복잡도: O(n*m)
- 쿼리의 범위 대로 부분 수열을 가져와서 오름차순 정렬한다.
  => 등차수열인 경우 자신 바로 앞의 숫자와 차이를 계산했을 때가 공차이다. (정렬해야만 이것이 성립한다.)
- 자신 바로 앞의 숫자와 차이를 계산해서 공차와 같은지 비교한다.
  => 하나라도 공차와 다른 경우가 있으면 등차수열이 아니다.
+) subArray := nums[l[i]:r[i]+1]을 하면 얕은 복사가 된다.
  => subArray가 nums의 원소를 참조하고 있으므로 subArray를 정렬했을 때 nums의 원소 위치도 같이 바뀐다.
  => 쿼리는 nums의 초기 상태를 기준으로 주어지므로 nums가 바뀌면 안된다.
  => subArray는 nums의 원소를 얕은 복사가 아닌 깊은 복사를 해야 한다.
  => subArray := append([]int{}, nums[l[i]:r[i]+1]...)
  => 위와 같이 하면 subArray는 nums의 원소를 참조하는 것이 아니라 그 값을 복사해서 사용한다.
  => subArray의 변경이 nums에 영향을 미치지 않는다.
*/
import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
    answer := []bool{}

    for i := 0; i < len(l); i++ {
        answer = append(answer, true)

        subArray := append([]int{}, nums[l[i]:r[i]+1]...)
        sort.Ints(subArray)

        diff := subArray[1] - subArray[0]

        for j := 1; j < len(subArray); j++ {
            if subArray[j] - subArray[j-1] != diff {
                answer[i] = false
                break
            }
        }
    }

    return answer
}
