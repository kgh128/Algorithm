/*
<큰 수 계산>
문제에서 0 이상의 두 정수가 문자열 a, b로 주어지는데,
해당 문자열의 길이는 최대 100,000이다.
따라서 a, b는 일반적인 int형의 범위를 넘어가는 수도 포함하고 있다.
=> int형 범위를 넘어가는 큰 수를 처리해야 한다.
*/


/*
1. math/big 라이브러리 이용
- 큰 수를 처리할 수 있도록 지원하는 라이브러리이다.

- big.Int: 큰 정수를 담을 수 있는 type
  => new()를 이용하여 할당
  => new(T)는 T형의 새로운 요소를 위해 zero화된 공간을 할당하고, type *T값인 해당 주소를 반환한다.
      (type T의 새롭게 할당된 zero 값에 대한 pointer를 반환)
      출처: https://ocwokocw.tistory.com/240

- (big.Int형 요소).SetString()
  - 첫번째 인자: 문자열
  - 두번째 인자: base (몇 진수인지)
  => 주어진 문자열을 주어진 base에 맞게 big.Int형으로 변환
  => 변환한 정수를 가리키는 포인터와 변환 성공 여부를 리턴한다.

- (big.Int형 요소).Add()
  - 더하고 싶은 두 정수를 받아서 더한 값을 리턴한다.
  - 큰 수 더하기가 가능하다.

- (big.Int형 요소).String()
  - 해당 정수를 문자열로 바꿔서 반환한다.
*/

import "math/big"

func solution(a string, b string) string {
    numA, _ := new(big.Int).SetString(a, 10)
    numB, _ := new(big.Int).SetString(b, 10)
    return new(big.Int).Add(numA, numB).String()
}



/*
2. 문자열로 주어진 큰 수를 한 자리씩 계산해서 이어붙이기
- 한 번에 큰 수를 정수로 바꿀 수는 없으니, 문자열을 돌면서 한 자리씩 정수로 바꿔서 계산한다.
*/

import (
    "strconv"
    "strings"
)

func solution(a string, b string) string {
    result := []string{}
    overflow := ""
    tens := 0

    // 1) a와 b의 길이를 똑같이 맞추고, 남는 숫자는 overflow에 저장
    // ex) a: "12345"
    //     b:   "345"
    // => a와 b는 모두 "345"로, overflow는 "12"로 설정
    if len(a) > len(b) {
        overflow = a[:len(a)-len(b)]
        a = a[len(a)-len(b):]
    } else {
        overflow = b[:len(b)-len(a)]
        b = b[len(b)-len(a):]
    }

    // 2) a와 b를 맨 뒤에서부터 돌면서 한 자리씩 덧셈 수행
    // -> 일의 자리부터 계산하는 효과
    // 만약 계산 결과가 10이 넘어가면 10의 자리수는 다음 계산으로 넘어간다.
    // 계산 결과의 일의 자리수는 result 배열에 추가된다.
    for i := len(a)-1; i >= 0; i-- {
        sum := int(a[i] - '0') + int(b[i] - '0') + tens
        tens = sum / 10
        result = append(result, strconv.Itoa(sum % 10))
    }

    // 3) 아까 길이가 넘어서 잘려나간 overflow 문자열도 2)와 동일하게 계산
    for i := len(overflow) - 1; i >= 0; i-- {
        sum := int(overflow[i] - '0') + tens
        tens = sum / 10
        result = append(result, strconv.Itoa(sum % 10))
    }

    // 4) 마지막 계산 결과가 10을 넘어서 10의 자릿수가 존재하는 경우, result에 추가
    if (tens > 0) {
        result = append(result, strconv.Itoa(tens))
    }

    // 5) 현재 result 배열은 일의 자릿수부터 집어넣은 상태이다. -> 숫자가 거꾸로 되어있다.
    // ex) 최종 결과는 "1234"가 나와야 하는데, 현재 result 배열에는 [4 3 2 1]로 되어있다.
    // 따라서 result 배열을 뒤집어서 올바른 결과로 만든다.
    for i := 0; i < len(result)/2; i++ {
        result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
    }

    // 6) string으로 변환하여 리턴
    return strings.Join(result, "")
}
