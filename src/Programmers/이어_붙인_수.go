/*
1. 문자열 이어붙이기
*/

import (
    "bytes"
    "strconv"
)

func solution(num_list []int) int {
    var odds bytes.Buffer
    var evens bytes.Buffer

    for _, num := range num_list {
        if (num % 2 == 1) {
            odds.WriteString(strconv.Itoa(num))
        } else {
            evens.WriteString(strconv.Itoa(num))
        }
    }

    odd_num, _ := strconv.Atoi(odds.String())
    even_num, _ := strconv.Atoi(evens.String())

    return odd_num + even_num 
}


/*
2. int 계산으로만 숫자 이어붙이기
현재까지 계산한 숫자를 sum, 다음에 이어붙일 숫자를 num이라고 할 때,
sum = sum * 10 + num을 하면 num이 가장 뒤에 이어붙여진 숫자가 sum이 된다.

ex) sum = 246, num = 8
sum * 10 + num = 2460 + 8 = 2468

=> 기존 숫자에 10을 곱함으로써 일의 자리에 공간을 하나 만들어준다.
=> 그리고 숫자를 더하면 일의 자리에 더해지므로 뒤에 이어붙여지는 효과가 나타난다.
*/

func solution(num_list []int) int {
    odds := 0
    evens := 0
    
    for _, num := range num_list {
        if (num % 2 == 1) {
            odds = odds * 10 + num
        } else {
            evens = evens * 10 + num
        }
    }

    return odds + evens
}
