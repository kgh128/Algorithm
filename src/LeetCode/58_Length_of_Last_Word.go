import "strings"

func lengthOfLastWord(s string) int {
    field := strings.Fields(s)
    return len(field[len(field)-1])
}
