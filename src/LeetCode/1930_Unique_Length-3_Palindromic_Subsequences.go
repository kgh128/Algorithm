/*
<알파벳의 가장 처음 인덱스와 끝 인덱스 저장하기>
- 시간복잡도: O(n), 공간복잡도: O(1)
- 첫번째 루프에서 문자열 s를 돌면서 해당 문자열에 속한 알파벳의 가장 처음 인덱스와 끝 인덱스를 map에 저장한다.
  => map의 value에 slice가 들어갈수도 있다.
  => 아직 map에 저장되지 않은 알파벳이면 현재 인덱스를 가장 처음 인덱스와 끝 인덱스로 설정하여 저장한다.
  => map에 이미 저장된 알파벳이면 현재 인덱스를 끝 인덱스에 저장한다.
- 두번째 루프를 돌면서 각 알파벳의 처음과 끝 인덱스 사이에 있는 알파벳 종류의 개수를 센다.
  => 그러면 그 사이에 있는 알파벳 종류의 개수가 현재 처음 인덱스와 끝 인덱스를 확인하고 있는 알파벳을 가지고 만든 unique palindromes의 개수가 된다.
  => 어차피 가장 앞과 뒤는 정해져 있고, 사이에 들어갈 알파벳 개수만 알면 되기 때문이다.
  => 중복된 알파벳은 한 번만 세야 하므로 사이에 있는 알파벳을 또다른 map에 저장한다.
  => 그러면 중복된 알파벳은 같은 key를 가지고 map에 접근하여 값을 true로 저장하기 때문에 나중에 map에 저장된 요소의 개수를 세면 알파벳 종류의 개수가 나온다.
*/
func countPalindromicSubsequence(s string) int {
    positions := make(map[byte][]int)  // 0: 가장 처음 인덱스, 1: 가장 끝 인덱스

    // 1. 각 알파벳의 가장 처음 인덱스와 끝 인덱스 찾기
    for i := 0; i < len(s); i++ {
        if _, exists := positions[s[i]]; !exists {
            positions[s[i]] = []int{i, i}
        } else {
            positions[s[i]][1] = i
        }
    }

    // 2. 각 알파벳의 가장 처음 인덱스와 끝 인덱스 사이에 있는 알파벳 개수 세기
    // 알파벳을 map을 이용해 세면 중복된 알파벳은 하나만 세진다.
    result := 0

    for _, pos := range positions {
        uniqueChars := make(map[byte]bool)

        for i := pos[0]+1; i < pos[1]; i++ {
            uniqueChars[s[i]] = true
        }

        result += len(uniqueChars) 
    }

    return result
}
