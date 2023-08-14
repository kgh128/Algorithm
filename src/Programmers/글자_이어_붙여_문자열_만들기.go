import "bytes"

func solution(my_string string, index_list []int) string {
    var ret bytes.Buffer

    for _, idx := range index_list {
        ret.WriteByte(my_string[idx])
    }

    return ret.String()
}
