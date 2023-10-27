// Two pointers
func isSubsequence(s string, t string) bool {
    sIdx, tIdx := 0, 0

    // 1. s와 t 중 하나가 끝에 도달할 때까지 반복문을 돈다.
    for sIdx < len(s) && tIdx < len(t) {
        if s[sIdx] == t[tIdx] {  // 2. 같은 문자이면 하나 찾은 것이므로 s의 인덱스를 하나 증가시켜 다음 문자를 찾게 한다.
            sIdx++
        }
        tIdx++  // 3. t의 인덱스는 항상 증가한다. -> s 문자열 순서 그대로 t 문자에 들어있어야 하기 때문에 이전 문자와 같지 않은 문자는 다음 문자에서도 의미가 없다.
    }

    // 4. s의 인덱스가 s의 길이와 같으면 s에 있는 모든 문자를 찾았다는 의미이므로 true가 리턴된다. (s 문자열이 끝나서 반복문 탈출)
    //    s의 인덱스가 s의 길이와 다르면 t에 있는 모든 문자를 검사했는데 s의 문자를 다 못찾았다는 의미이므로 false가 리턴된다. (t 문자열이 끝나서 반복문 탈출)
    return sIdx == len(s)
}
