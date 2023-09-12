// 1. result 배열 길이를 a, b 배열 중 더 긴 길이로 미리 설정
func addBinary(a string, b string) string {
    var result []byte
    sum := 0
    aIdx := len(a) - 1
    bIdx := len(b) - 1

    if len(a) > len(b) {
        result = make([]byte, len(a) + 1)
    } else {
        result = make([]byte, len(b) + 1)
    }

    for i := len(result) - 1; i >= 0; i-- {
        if aIdx >= 0 {
            sum += int(a[aIdx] - '0')
            aIdx--
        }
        if bIdx >= 0 {
            sum += int(b[bIdx] - '0')
            bIdx--
        }
        result[i] = byte(sum % 2 + '0')
        sum /= 2
    }

    if result[0] == '0' {
        return string(result[1:])
    } else {
        return string(result)
    }
}


// 2. result 배열을 append로 늘려나감.
// -> 새로 추가하는 값을 append의 첫번째 인자로 주어서 result 배열 앞에 값이 추가될 수 있도록 함.
func addBinary(a string, b string) string {
    var result []byte
    sum := 0
    aIdx := len(a) - 1
    bIdx := len(b) - 1

    for aIdx >= 0 || bIdx >= 0 {
        if aIdx >= 0 {
            sum += int(a[aIdx] - '0')
            aIdx--
        }
        if bIdx >= 0 {
            sum += int(b[bIdx] - '0')
            bIdx--
        }

        result = append([]byte{byte(sum % 2 + '0')}, result...)
        sum /= 2
    }

    if sum == 1 {
        return string(append([]byte{'1'}, result...))
    } else {
        return string(result)
    }
}
