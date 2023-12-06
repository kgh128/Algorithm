/*
1. 반복문으로 풀기
- 시간복잡도: O(n), 공간복잡도: O(1)
*/
func totalMoney(n int) int {
    total, daily, monday := 0, 0, 0

    for i := 0; i < n; i++ {
        if i % 7 == 0 {
            monday++
            daily = monday
        } else {
            daily++
        }
        
        total += daily
    }

    return total
}


/*
2. 수학 계산으로 풀기
- 시간복잡도: O(1), 공간복잡도: O(1)
*/
func totalMoney(n int) int {
    week := n / 7
    day := n % 7

    fullWeek := 28 * week + 7 * (week - 1) * week / 2
    lastWeek := day * week + day * (day + 1) / 2

    return fullWeek + lastWeek
}
