/*
1. 나머지 연산으로 홀수 판별
- 시간복잡도: O(n), 공간복잡도: O(1)
- 일의 자리부터 반복문으로 탐색하여 홀수가 나오면 리턴
*/
func largestOddNumber(num string) string {
  for i := len(num) - 1; i >= 0; i-- {
      if int(num[i] - '0') % 2 == 1 {
          return num[:i+1]
      }
  }
  return ""
}


/*
2. 비트 연산으로 홀수 판별
- 시간복잡도: O(n), 공간복잡도: O(1)
- 모든 홀수는 첫번째 비트가 1이다. 그래야만 유일하게 홀수인 2^0(=1)을 얻을 수 있다.
- 나머지 비트는 2^1(=2)부터 시작되는 2의 제곱수를 의미하므로 짝수만을 얻을 수 있다.
- 일의 자리부터 반복문으로 탐색하여 홀수가 나오면 리턴
*/
func largestOddNumber(num string) string {
  idx := len(num) - 1

  for idx >= 0 && (num[idx] - '0') & 1 == 0 {
    idx--
  }

  return num[:idx+1]
}
