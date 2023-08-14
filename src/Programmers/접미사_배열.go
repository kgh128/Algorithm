import "sort"

func solution(my_string string) []string {
    var suffix []string

    for idx, _ := range my_string {
        suffix = append(suffix, my_string[idx:])
    }

    sort.Strings(suffix)

    return suffix
}
