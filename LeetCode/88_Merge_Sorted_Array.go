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
