/*
<정렬 + 투 포인터로 풀기>
- 시간복잡도: O(n) -> 최악의 경우 len(g)와 len(s) 중 작은 것만큼 반복문을 돈다.
- 공간복잡도: O(1)
- 슬라이스 정렬 방법: slices.Sort()
  => 지금까지 sort를 import하여 sort.Ints()와 같은 함수를 썼는데 이런 방법도 있었다.
  => []int를 정렬할 때는 sort.Ints()보다 더 빠른 것 같다.
  => sort.Ints() 썼을 때는 21ms~27ms가 나왔는데, slices.Sort()를 쓰니까 18ms가 나왔다.
*/

func findContentChildren(g []int, s []int) int {
    maxNum := 0
    i, j := 0, 0

    slices.Sort(g)
    slices.Sort(s)

    for i < len(g) && j < len(s) {
        if s[j] >= g[i] {
            maxNum++
            i++
        }
        j++
    }

    return maxNum
}
