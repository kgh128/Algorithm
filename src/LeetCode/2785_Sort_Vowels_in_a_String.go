/*
<모음만 따로 빼서 정렬하고 다시 집어넣기>
- 시간복잡도: O(n), 공간복잡도: O(n)
- 첫번째 루프에서 모음만 따로 int 슬라이스에 저장하기
  => 바이트 슬라이스는 sort 패키지에 미리 정의된 정렬 함수가 존재하지 않으므로 byte를 int로 바꾸어 저장한다.
  => 그리고 int 슬라이스를 sort.Ints() 함수로 내림차순 정렬하였다.
- 주어진 string 변수 s를 byte 슬라이스를 변환하였다. -> 각 문자의 값을 바꿀 수 있도록 (string 변수에서는 각 문자에 값을 할당하지 못한다.)
- 두번째 루프를 돌면서 s의 byte 슬라이스에 모음이 나오면 정렬해둔 모음 슬라이스에서 값을 꺼내서 해당 문자에 할당한다.
  => 이때 정렬해둔 모음은 int형이므로 byte형으로 바꾸어서 할당한다.
*/
import "sort"

func isVowel(c byte) bool {
    return c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U'
}

func sortVowels(s string) string {
    vowels := []int{}

    for i := 0; i < len(s); i++ {
        if isVowel(s[i]) {
            vowels = append(vowels, int(s[i]))
        }
    }

    sort.Ints(vowels)

    result := []byte(s)
    vowelIndex := 0

    for i := 0; i < len(s); i++ {
        if isVowel(s[i]) {
            result[i] = byte(vowels[vowelIndex])
            vowelIndex++
        }
    }

    return string(result)
}
