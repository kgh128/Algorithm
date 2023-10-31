// 1. 반복문 이용
// - n: columnNumber
// - 시간복잡도: O(log26n), 공간복잡도: O(log26n)
func convertToTitle(columnNumber int) string {
    result := []byte{}

    for columnNumber != 0 {
        columnNumber--
        result = append([]byte{'A' + byte(columnNumber % 26)}, result...)
        columnNumber /= 26
    }

    return string(result)
}


// 2. 재귀호출 이용
// - n: columnNumber
// - 시간복잡도: O(log26n), 공간복잡도: O(log26n)
