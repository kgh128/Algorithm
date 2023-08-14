import "strconv"

func solution(intStrs []string, k int, s int, l int) []int {
    var ret []int

    for _, intStr := range intStrs {
        num, _ := strconv.Atoi(intStr[s:s+l])

        if num > k {
            ret = append(ret, num)
        }
    }

    return ret
}
