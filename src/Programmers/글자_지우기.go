/*
1. 지우는 글자들 사이의 문자열을 차례대로 이어붙이기
ex) 0 1 2 3 4 5 6의 인덱스를 가지고 있는 문자열
=> 지우는 인덱스: 4, 1, 5 -(정렬)-> 1, 4, 5  // 정렬을 해서 문자열 앞에서부터 차례대로 문자를 지우자.
- 처음 구간: str[:1] => str[0]이 버퍼에 추가된다.
- 중간 구간: str[2:4], str[5:5] => tr[2:4]가 버퍼에 추가된다.
- 끝 구간: str[6:] => str[6]이 버퍼에 추가된다.
=> 결과: 0 2 3 6
*/

import(
    "sort"
    "bytes"
)

func solution(myStr string, rmvIdx []int) string {
    sort.Ints(rmvIdx)

    var retStr bytes.Buffer
    var start, end = rmvIdx[0], rmvIdx[len(rmvIdx)-1]

    // 1. 처음 구간
    retStr.WriteString(myStr[:start])

    // 2. 중간 구간
    for i := 0; i < len(rmvIdx) - 1; i++ {
        retStr.WriteString(myStr[rmvIdx[i]+1:rmvIdx[i+1]])
    }

    // 3. 끝 구간
    retStr.WriteString(myStr[end+1:])

    return retStr.String()
}


/*
2. strings.ReplaceAll() 이용하기
- 삭제 인덱스 배열을 돌면서 해당 인덱스의 문자를 공백으로 바꾼다. (문자 지우기)
- strings.ReplaceAll()을 이용하여 최종 문자열의 모든 공백을 없앤다. (이어붙이기) 
*/

import "strings"

func solution(myStr string, indices []int) string {
    retStr := []byte(myStr)
    
    for _, idx := range indices {
        retStr[idx] = ' '
    }
    
    return strings.ReplaceAll(string(retStr), " ", "")
}
