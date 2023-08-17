func solution(arr []int, n int) []int {
    var start int

    if len(arr) % 2 == 1 {
        start = 0
    } else {
        start = 1
    }

    for i := start; i < len(arr); i += 2 {
        arr[i] += n
    }

    return arr
}
