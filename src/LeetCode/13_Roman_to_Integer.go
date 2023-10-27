/*
<로마자 문자열을 뒤에서부터 순회하기>
1. map에 로마자-숫자 쌍을 저장
2. 현재 로마자를 curr, 현재 로마자 직전의 로마자를 prev로 설정
3. 로마자 문자열을 뒤에서부터 순회하므로 일반적으로 curr > prev이다.
   로마자는 앞으로 갈수록 큰 숫자가 존재하므로, 뒤에 있는 로마자가 더 작은 숫자.
   ex) XXVII
   => 이 경우 total에 curr을 더한다.
4. 단, 특수한 경우 더 작은 숫자가 더 큰 숫자보다 앞에 존재하는 경우가 있다.
   이는 더 큰 숫자에서 더 작은 숫자를 뺀 숫자를 의미한다.
   ex) IV = 4  ->  I가 V보다 더 작은 숫자인데 더 앞에 있다.
       그리고 5(V) - 1(I) = 4를 의미한다.
   => 이때는 curr = 1, prev = 5일텐데, curr < prev
   => prev는 이전 루프에서 이미 total에 더해져있다.
   => 따라서 curr를 total에서 빼주면 원하는 숫자가 total에 들어간다.
   ex) IV = 4 -> total + 5 - 1 = total + 4
5. 모든 연산을 끝내면 prev에 curr를 넣고 다음 루프로 넘어간다.
*/

func romanToInt(s string) int {
    var total, prev, curr int
    convert := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    for i := len(s) - 1; i >= 0; i-- {
        curr = convert[s[i]]

        if curr < prev {
            total -= curr
        } else {
            total += curr
        }

        prev = curr
    }

    return total
}
