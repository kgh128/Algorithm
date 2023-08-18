import "sort"

func solution(num_list []int) []int {
    sort.Ints(num_list)
    return num_list[5:]
}
