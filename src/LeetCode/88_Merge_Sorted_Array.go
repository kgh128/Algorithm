// 1. sort 라이브러리 사용
// - nums1과 nums2를 이어붙인 후 sort.Inst()로 정렬
// - 현재는 2ms 나온다. 시간 복잡도 자체는 O((n+m)log(n+m))
// - 이어붙일 때가 (n+m), sort.Ints()로 정렬할 때가 log(n+m)
import "sort"

func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1 = append(nums1[:m], nums2...)
    sort.Ints(nums1)
}


// 2. 직접 merge sort
// - 0ms 나온다. 시간 복잡도는 O(m+n)
// - 기존에 주어진 m과 n만 이용하여 merge sort를 함. (투 포인터)
//   => m과 n을 인덱스로 사용하기 위해 처음에 m--, n--를 해줌. 
// - num1이 결과 배열이므로 num1은 (m + n - 1) 길이의 배열로 이미 설정되어 있음.
//   => num1[m:]은 0으로 차있는 slice임.
// - 따라서 merge sort를 할 때 num1[m]과 num2[n] 중에서 큰 것을 num1의 뒤에서부터 집어넣음.
//   => 뒤에서부터 집어넣을 때의 인덱스는 (m + n + 1)이다.
//   => 남은 길이: m + 1 + n + 1 = m + n + 2
//   => 가장 마지막 인덱스 = 남은 길이 - 1 = m + n + 1
// - 둘 중에서 작은 것을 num1의 앞에서부터 집어넣는 방식을 사용하면 num1[:m]에 있는 값들이 사라지기 때문에 0으로 비워져있는 뒤에서부터 넣는 방식을 채택.
// - for문의 조건이 n >= 0인 이유는 num2에 남은 숫자가 없어서 n < 0이 되면 남은 것은 num1에 있는 숫자들인데 이는 이미 num1에 정렬되어서 넣어져 있는 상태이기 때문에 건드릴 필요가 없음.
// - 뒤에서부터 넣는 방식을 채택하여 앞에 있는 num1의 숫자를 건드리지 않기 때문에 가능한 방식임.
func merge(nums1 []int, m int, nums2 []int, n int)  {
    m--
    n--

    for n >= 0 {
        if m >= 0 && nums1[m] > nums2[n] { // nums1의 숫자가 남고, 그 숫자가 nums2의 숫자보다 클 때
            nums1[m + n + 1] = nums1[m]
            m--
        } else { // nums1의 숫자가 안남거나(nums2의 숫자만 남아있거나), nums1의 숫자보다 nums2의 숫자가 더 클 때
            nums1[m + n + 1] = nums2[n]
            n--
        }
    }
}
