// 1. 반복문으로 풀기 (비트 연산X)
// 시간복잡도: O(logn), 공간복잡도: O(1)
func isPowerOfFour(n int) bool {
    comp := 1

    for comp <= n {
        if n == comp {
            return true
        }
        
        comp *= 4
    } 

    return false
}

// 2. 비트 연산으로 풀기 (반복문 사용X)
// 시간복잡도: O(1), 공간복잡도: O(1)
// 1) 2의 배수인지 확인 => (n & (n-1) == 0)
// - 4의 배수도 2의 배수이므로 일단 2의 배수인지를 확인하여 거른다.
// - 2의 배수는 전체 비트에서 1이 한 개만 존재한다.
// - 따라서 2의 배수에서 1을 빼면 원래 자리에 있던 1은 0이 되고, 그 뒤의 비트들이 모두 1이 된다.
// - 따라서 n과 n-1을 & 연산을 하면 결과값이 0이 된다.
// 2) 2의 배수 중에서 4의 배수인지 확인 => (n & mask == n)
// - 4의 배수는 비트에 딱 하나 존재하는 1이 홀수 자리에 존재한다. (나머지 2의 배수들은 짝수 자리에 존재)
// ex) 4 = 0000 0100, 16 = 0001 0000
// - 따라서 모든 홀수 자리에 1이 존재하는 mask 값과 & 연산을 하면 원래 존재하는 1만 살아남아서 결과값이 자기 자신이 된다.
// - 최댓값이 2^31 - 1 = 0111 1111 1111 1111 1111 1111 1111 1111
// - 따라서 최댓값까지 커버해야 하는 mask 값은 0x55555555 = 0101 0101 0101 0101 0101 0101 0101 0101
func isPowerOfFour(n int) bool {
  mask := 0x55555555
  return (n > 0) && (n & (n-1) == 0) && (n & mask == n)
}
