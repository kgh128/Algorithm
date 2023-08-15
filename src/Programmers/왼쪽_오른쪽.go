func solution(str_list []string) []string {
    for idx, str := range str_list {
        if (str == "l") {
            return str_list[:idx]
        } else if (str == "r") {
            return str_list[idx+1:]
        }
    }
    return []string{}
}
