/*
<논리 연산>
- 논리합: OR, ∨, ||
- 논리곱: AND, ∧, &&
*/

func solution(x1 bool, x2 bool, x3 bool, x4 bool) bool {
    return (x1 || x2) && (x3 || x4)
}
