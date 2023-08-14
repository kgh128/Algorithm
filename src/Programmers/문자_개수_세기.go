/*
1. 대문자의 인덱스
=> 시작점 'A'로부터 얼만큼 떨어져 있는가이다.

2. 소문자의 인덱스
=> 소문자 전에 이미 대문자 26개가 존재하므로 소문자의 인덱스는 26부터 시작한다.
=> 시작점 'a'로부터 얼만큼 떨어져 있는가를 26에 더해준다.
*/

func solution(my_string string) []int {
    count := make([]int, 52)

    for _, char := range my_string {        
        if char <= 'Z' {
            count[int(char-'A')]++
        } else {
            count[26 + int(char-'a')]++
        }
    }

    return count
}
