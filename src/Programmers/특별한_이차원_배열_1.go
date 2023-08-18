/*
<이차원 슬라이스 만들기>
arr := make([][]int, n)   // 이차원 슬라이스 생성
arr[i] := make([]int, n)  // 이차원 슬라이스의 각 원소(행)에 일차원 슬라이스를 할당
*/

func solution(n int) [][]int {
    arr := make([][]int, n)
    
    for i := 0; i < n; i++ {
        arr[i] = make([]int, n)
        arr[i][i] = 1
    }
    
    return arr
}
