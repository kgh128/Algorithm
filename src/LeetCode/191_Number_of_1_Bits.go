/*
1. 비트 연산자 이용하여 직접 세기
- 시간복잡도: O(n), 공간복잡도: O(1)
- 가장 오른쪽에 있는 비트가 0인지 1인지 확인해서 1이면 count 증가
- 오른쪽 쉬프트 연산자를 이용해서 num의 비트를 오른쪽으로 한 칸씩 옮기기
*/
func hammingWeight(num uint32) int {
    if num == 0 {
        return 0
    }

    var count int = 0
    var mask uint32 = 1

    for num > 0 {
        if (num & mask) == 1 {
            count++
        }
        num = num >> 1
    }

    return count
}


/*
2. math/bits 패키지 이용
- bits.OnesCount() 함수를 사용하면 비트로 표현했을 때 1의 개수를 세준다.
- OnesCount32() 함수를 사용한 이유는 num의 타입이 uint32이기 때문이다.
  => OnesCount(): 입력으로 uint 변수를 받는다.
     이 함수를 쓰고 싶으면 bits.OnesCount(uint(num))으로 호출하면 된다. (num을 uint 타입으로 바꿔서 전달)
  => OnesCount32(): 입력으로 uint32 변수를 받는다.
*/
import "math/bits"

func hammingWeight(num uint32) int {
    return bits.OnesCount32(num)
}
