/*
1. 열 추가: 각 행마다 (행의 개수) - (열의 개수) 길이의 0으로 채워진 int slice를 append
ex)
1 2          1 2 0
3 4    =>    3 4 0
5 6          5 6 0

2. 행 추가: 전체 배열에 (열의 개수) 길이의 0으로 채워진 int slice를 (행의 개수) - (열의 개수)만큼 append
ex)
1 2 3 4          1 2 3 4
4 5 6 7    =>    4 5 6 7      
8 9 1 2          8 9 1 2
                 0 0 0 0
*/


func solution(arr [][]int) [][]int {
    if len(arr) > len(arr[0]) {  // 1. 열 추가
        addLength := len(arr) - len(arr[0])

        for i, row := range arr {
            arr[i] = append(row, make([]int, addLength)...)
        }
    } else {  // 2. 행 추가
        addLength := len(arr[0]) - len(arr)

        for i := 0; i < addLength; i++ {
            arr = append(arr, make([]int, len(arr[0])))
        }
    }

    return arr
}
