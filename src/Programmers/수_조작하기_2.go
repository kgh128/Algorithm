import "bytes"

func solution(numLog []int) string {
    var control bytes.Buffer

    for i := 1; i < len(numLog); i++ {
        switch numLog[i] - numLog[i-1] {
            case 1:
                control.WriteByte('w')
            case -1:
                control.WriteByte('s')
            case 10:
                control.WriteByte('d')
            case -10:
                control.WriteByte('a')
        }
    }

    return control.String()
}
