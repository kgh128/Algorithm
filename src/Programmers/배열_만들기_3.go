/*
<Go에서 slice 이어붙이기>
- append 두번쨰 인자(이어붙이고 싶은 slice에 ...을 붙여줘야 함.
=> ...은 가변 인자를 의미함. (함수에 임의 개수의 인자를 넘길 수 있음.)
=> slice는 자료형 인자 여러 개의 묶음인데, append에서 ...을 통해 slice에 있는 인자 여러 개를 넘길 수 있음.
- 링크 1: https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
- 링크 2: https://go.dev/ref/spec#Appending_and_copying_slices
- 링크 3: https://go.dev/ref/spec#Passing_arguments_to_..._parameters

ex)
append(s S, x ...T) S  // T is the element type of S

s0 := []int{0, 0}
s1 := append(s0, 2)        // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)  // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)    // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
*/

// 1. 반복문 이용
func solution(arr []int, intervals [][]int) []int {
    result := []int{}

    for _, interval := range intervals {
        result = append(result, arr[interval[0]:interval[1]+1]...)
    }

    return result
}

// 2. 어차피 두 구간 밖에 없으니까 한 번에 이어붙이기
func solution(arr []int, intervals [][]int) []int {
    var a1, b1, a2, b2 = intervals[0][0], intervals[0][1], intervals[1][0], intervals[1][1]
    
    return append(arr[a1:b1+1], arr[a2:b2+1]...)
}
