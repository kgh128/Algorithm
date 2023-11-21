/*
1. 각 쓰레기마다 배열이 존재하고, 해당 배열에 구간합 저장
- 시간복잡도: O(n*m) -> n은 집의 개수, m은 해당 집의 쓰레기 개수
- 공간복잡도: O(n) -> 각 쓰레기마다 배열이 존재
- 각 쓰레기의 구간합을 저장하는 배열은 해당 집에 해당 쓰레기가 있는지와는 상관없이 무조건 travel[i] 값을 기본적으로 더하고 시작한다.
  => 현재 집에는 해당 쓰레기가 없지만, 다음에 나타날 집에는 해당 쓰레기가 있을 수 있으므로 이와 같은 방식을 취한다.
- 각 쓰레기의 마지막 위치도 계속 추적한다. 각 쓰레기의 마지막 위치에 해당하는 구간합을 더해서 결과로 제출한다.
*/
func garbageCollection(garbage []string, travel []int) int {
    lastHouse := len(garbage) - 1
    lastM, lastP, lastG := 0, 0, 0
    metal, paper, glass := []int{0}, []int{0}, []int{0}

    for i := 0; i <= lastHouse; i++ {
        for _, garbageType := range garbage[i] {
            if garbageType == rune('M') {
                lastM = i
                metal[i]++
            } else if garbageType == rune('P') {
                lastP = i
                paper[i]++
            } else {
                lastG = i
                glass[i]++
            }
        }

        if i < lastHouse {
            metal = append(metal, metal[i] + travel[i])
            paper = append(paper, paper[i] + travel[i])
            glass = append(glass, glass[i] + travel[i])
        }
    }

    return metal[lastM] + paper[lastP] + glass[lastG]
}


/*
2. 각 쓰레기마다 변수가 존재하고, 해당 변수에 이동 시간을 저장 + 쓰레기 수거 소요 시간은 따로 관리
- 시간복잡도: O(n*m) -> n은 집의 개수, m은 해당 집의 쓰레기 개수
- 공간복잡도: O(1) -> 각 쓰레기마다 배열이 존재하는 것이 아니라 변수가 존재
- 각 집에서의 쓰레기 수거 소요 시간은 해당 집의 쓰레기 개수와 같다.
- currTravelTime 변수에 현재 위치까지의 이동 시간을 저장한다. => 구간합 이용
- 해당 집에서 나온 쓰레기의 이동 시간은 currTravelTime이 된다. 그 쓰레기를 수거하러 해당 위치까지 이동해야 하기 때문이다.
  => 쓰레기의 모든 위치의 이동 시간을 배열로 관리할 필요 없이, 최종 이동 시간만 변수로 관리하면서 계속 갱신해주면 된다.
  => 공간을 더 적게 사용한다.
*/
func garbageCollection(garbage []string, travel []int) int {
    collectTime := 0
    travelM, travelP, travelG := 0, 0, 0
    currTravelTime := 0

    for i := 0; i < len(garbage); i++ {
        if i > 0 {
            currTravelTime += travel[i-1]
        }

        collectTime += len(garbage[i])

        for _, garbageType := range garbage[i] {
            if garbageType == rune('M') {
                travelM = currTravelTime
            } else if garbageType == rune('P') {
                travelP = currTravelTime
            } else {
                travelG = currTravelTime
            }
        }
    }

    return collectTime + travelM + travelP + travelG
}
