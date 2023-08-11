/*
<make를 이용하여 Slice 생성>
첫번째 안자: Slice type
두번째 인자: length
세번째 인자: capacity (생략하면 capacity = length)
-> length가 상수가 아닌 변수인 배열을 만들 수 있음.

var arr [3]int
위와 같이 배열을 선언할 경우, length에 예시의 3처럼 상수만 올 수 있음.
var arr [length]int -> 컴파일 에러
*/

func solution(start int, end int) []int {
    length := start - end + 1
    arr := make([]int, length)
    
    for i := 0; i < length; i++ {
        arr[i] = start
        start--
    }
    
    return arr
}
