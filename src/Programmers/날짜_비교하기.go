// 1. 반복문으로 원소 비교
func solution(date1 []int, date2 []int) int {
    for i, _ := range date1 {
        if date1[i] < date2[i] {
            return 1
        } else if date1[i] > date2[i] {
            return 0
        }
    }
    return 0
}


// 2. 숫자로 만들어서 크기 비교
// date[2]: 두 자리수 날짜 -> 그냥 더하기
// date[1]: 두 자리수 월   -> 뒤에 두 자리수 날짜가 있으므로 100 곱해서 더하기
// date[0]: 네 자리수 연도 -> 뒤에 총 네 자리수 월과 날짜가 있으므로 10000 곱해서 더하기
func solution(date1 []int, date2 []int) int {
    d1 := 10000*date1[0] + 100 *date1[1] + date1[2]
    d2 := 10000*date2[0] + 100 *date2[1] + date2[2]
    
    if d1 < d2 {
        return 1
    }
    return 0
}
