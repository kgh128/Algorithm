/*
<시작 인덱스>
- 문제에서는 시작 인덱스를 1로 정의하였는데, 실제 배열 인덱스는 0부터 시작한다.
  (실제 인덱스는 문제 인덱스 - 1)
- 따라서 문제 인덱스의 홀짝과 실제 인덱스의 홀짝은 서로 바뀐다.
- 문제의 홀수 인덱스(1, 3, 5, ...) => 실제 짝수 인덱스(0, 2, 4, ...)
- 문제의 짝수 인덱스(2, 4, 6, ...) => 실제 홀수 인덱스(1, 3, 5, ...)
*/

func solution(num_list []int) int {
    var odd_sum, even_sum = 0, 0

    for idx, num := range num_list {
        if idx % 2 == 0 {
            odd_sum += num
        } else {
            even_sum += num
        }
    }

    if (odd_sum > even_sum) {
        return odd_sum
    } else {
        return even_sum
    }
}
