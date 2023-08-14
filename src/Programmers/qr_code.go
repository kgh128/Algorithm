/*
1. code의 문자를 반복문으로 다 돌면서 조건에 맞는 인덱스 찾기
*/

func solution(q int, r int, code string) string {
    str := []rune{}

    for idx, char := range code {  // len(code)만큼 반복문을 돎.
        if idx % q == r {
            str = append(str, char)
        }
    }

    return string(str)
}


/*
2. 인덱스를 r부터 시작해서 q씩 늘리기
- i: q로 나눴을 때 나머지가 r인 숫자
=> i = q * (몫) + r
ex) i = q * 0 + r, i = q * 1 + r, i = q * 2 + r, ...
=> 따라서 i를 r부터 시작해서 q씩 늘리면 위의 식을 만족하는 i를 얻을 수 있음.
   (q에 곱해지는 몫이 하나씩 늘어남.)
*/

func solution(q int, r int, code string) string {
    str := []byte{}
    
    for i := r; i < len(code); i += q {  // len(code)/q만큼 반복문을 돎.
        str = append(str, code[i])
    }
    
    return string(str)
}
