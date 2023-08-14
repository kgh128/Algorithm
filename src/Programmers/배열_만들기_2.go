import "strconv"

func solution(l int, r int) []int {
    var result []int

    for i := l; i <= r; i++ {
        is_zero_or_five := true

        for _, num := range strconv.Itoa(i) {
            if num != '0' && num != '5' {
                is_zero_or_five = false
                break
            }
        }

        if is_zero_or_five {
            result = append(result, i)
        }
    }

    if len(result) == 0 {
        result = append(result, -1)
    }

    return result
}
