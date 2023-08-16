import "strings"

func solution(myString string) []int {
    lenArr := []int{}
    strArr := strings.Split(myString, "x")

    for _, str := range strArr {
        lenArr = append(lenArr, len(str))
    }

    return lenArr
}
