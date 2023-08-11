/*
<문자로 표현된 숫자를 실제 숫자로 바꾸기>
ex) '5' -> 5 // rune -> int

1. 문자로 표현된 숫자('5'): rune 타입
2. 실제 숫자(5): int 타입

var num rune = '5'
-> int(num)은 '5'의 아스키코드 값인 53
-> int(num) - '0' 또는 int(num - '0')이 실제 숫자값인 5

=> '0'의 아스키코드는 48이므로 이를 빼주는 것은 해당 문자가 '0'부터 몇번째 이후에 있는지를 의미한다.
    0부터 5번째 뒤에 있는 숫자는 결국 5이므로 이를 통해 원하는 숫자값을 얻을 수 있다. 

*/

func solution(num_str string) int {
    sum := 0
    
    for _, num := range num_str {
        sum += (int(num) - '0')
    }

    return sum
}
