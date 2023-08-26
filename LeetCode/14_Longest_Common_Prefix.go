/*
1. 가장 첫번째 문자열을 기준으로 모든 문자열과 비교하기
- prefix의 최대 길이는 가장 짧은 문자열의 길이
  => 바깥쪽 for문은 가장 짧은 문자열의 길이만큼만 돈다.
- 모든 문자열의 문자를 하나씩 비교하면서 하나라도 문자가 다른 문자열이 나오면 지금까지 구한 prefix를 반환한다.
- 모든 문자열의 현재 순회하고 있는 문자가 같으면 해당 문자는 prefix에 추가한다.
*/
func longestCommonPrefix(strs []string) string {
    prefix := []byte{}
    minLen := 200

    for _, str := range strs {
        if len(str) < minLen {
            minLen = len(str)
        }
    }

    for i := 0; i < minLen; i++ {
        char := strs[0][i]
        
        for j := 1; j < len(strs); j++ {
            if strs[j][i] != char {
                return string(prefix)
            }
        }

        prefix = append(prefix, char)
    }

    return string(prefix)
}


/*
2. 문자열을 정렬한 후 가장 첫번째 문자열과 가장 마지막 문자열만 비교하기
- 문자열을 정렬하면 가장 첫번째 문자열과 가장 마지막 문자열이 가장 짧은 공통 접두사를 가지고 있다.
ex) googleaaa googleaaab googleaaabc googleaac googleb => google
ex) googlebcd googlebce googlebcef googlebcf googlebdd => goolgeb
=> 오름차순으로 정렬되므로 첫번째 문자열과 공통된 부분이 많을수록 정렬되었을 때 앞에 위치할 수 밖에 없다.
   (뒤에 있는 문자열은 첫번째 문자열과의 공통 접두사 뒤에 문자가 붙은 문자열이니까)
=> 따라서 첫번째 문자열과 마지막 문자열의 공통 접두사가 모든 문자열들의 공통 접두사이다.
   (첫번째 문자열과 마지막 문자열이 가장 겹치는 부분이 적다.)
*/
import "sort"

func longestCommonPrefix(strs []string) string {
    sort.Strings(strs)

    l := len(strs)

    for i, _ := range strs[0] {
        if strs[0][i] != strs[l-1][i] {
            return strs[0][:i] // 다른 문자가 감지되는 순간 그 전까지만 잘라서 리턴
        }
    }

    return strs[0]
}
