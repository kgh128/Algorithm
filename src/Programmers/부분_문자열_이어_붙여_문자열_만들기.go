import "bytes"

func solution(myStr []string, parts [][]int) string {
    var retStr bytes.Buffer

    for idx, part := range parts {
        var s, e = part[0], part[1]
        retStr.WriteString(myStr[idx][s:e+1])
    }

    return retStr.String()
}
