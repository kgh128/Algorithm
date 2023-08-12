import "bytes"


func solution(str1 string, str2 string) string {
    var buffer bytes.Buffer

    for i := 0; i < len(str1); i++ {
        buffer.WriteByte(str1[i])
        buffer.WriteByte(str2[i])
    }

    return buffer.String()
}
