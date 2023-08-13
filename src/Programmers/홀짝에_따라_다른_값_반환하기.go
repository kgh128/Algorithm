/*
<홀수, 짝수 간격>
홀수와 홀수 사이의 간격도 2이고, 짝수와 짝수 사이의 간격도 2이다.
따라서 반복문에서 i의 값이 2씩 줄어들게 하여 그 다음 홀수나 짝수를 구하였다.
*/

func solution(n int) int {
    sum := 0

    if n%2 == 1 {
        for i := n; i > 0; i -= 2 {
            sum += i
        }
    } else {
        for i := n; i > 0; i -= 2 {
            sum += (i * i)
        }
    }

    return sum
}
