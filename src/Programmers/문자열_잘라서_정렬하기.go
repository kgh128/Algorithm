/*
<strings.FiledsFunc() 이용하기>
- 첫번쨰 인자: 자르고 싶은 문자열
- 두번째 인자: 문자열을 자를 기준이 될 함수 (함수 구현부를 바로 인자에 넣을 수 있다.)
  => 이 함수는 사용자가 직접 정의하는 유니코드 검사 함수로, 해당 함수를 만족하는 특정 유니코드 값을 기준으로 문자열을 쪼갠다.
ex) strings.FieldsFunc(myString, func(r rune) bool {
        return r == 'x'
    })
    이렇게 코드를 짜면 'x'를 기준으로 myString을 자른다.
    1) 'x'가 연속으로 나와도 결과물인 문자열 슬라이스에 빈 문자열이 들어가지 않는다.
    => strings.split()에서는 자르는 기준이 될 문자열에 "x"만 넣으면 "x"가 연속으로 나왔을 때,
      "x" 하나를 기준으로 자르므로 "x"들 사이의 빈 문자열이 결과 문자열 슬라이스에 들어간다.
    => 그를 방지하려면 "x+"라는 정규식을 기준으로 문자열을 잘라야한다.(strings가 아닌 regexp 라이브러리 사용해야 함.)
    2) 원본 문자열의 맨 앞이나 맨 뒤가 "x"여도 결과물에 빈 문자열이 들어가지 않는다.
    => strings.split()나 regexp 라이브러리를 이용하여 정규식을 정의하고, 해당 정규식을 기준으로 split하면 위의 경우에 빈 문자열이 결과물에 들어간다.
    => 맨 앞이나 맨 뒤에 "x"가 있는 문자열을 "x"나 "x+"를 기준으로 자르면 "x" 앞이나 뒤에 아무것도 없으니 잘랐을 때 빈 문자열이 나오기 때문이다.
*/

import (
    "strings"
    "sort"
)

func solution(myString string) []string {
    strArr := strings.FieldsFunc(myString, func(r rune) bool {
        return r == 'x'
    })
    sort.Strings(strArr)
    return strArr
}
