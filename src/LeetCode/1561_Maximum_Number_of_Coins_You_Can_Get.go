/*
<정렬하고 현재 배열에서 두번째로 큰 값 가져가기>
- 시간복잡도: O(nlogn), 공간복잡도: O(1)
- piles를 정렬하면 가장 작은 값부터 가장 큰 값의까지의 위치를 파악할 수 있다.
- 코인 더미 3개를 고르는 기준: [가장 큰 값, 두번째로 큰 값, 가장 작은 값]
  => 이렇게 고르면 항상 내가 두번째로 큰 값을 가져간다.
*/
import "sort"

func maxCoins(piles []int) int {
    result := 0 

    sort.Ints(piles)

    for len(piles) > 0 {
        secondMax := len(piles) - 2
        result += piles[secondMax]
        piles = piles[1:secondMax]
    }    

    return result
}
