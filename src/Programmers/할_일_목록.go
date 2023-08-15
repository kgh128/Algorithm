func solution(todo_list []string, finished []bool) []string {
    unfinished := []string{}

    for idx, isFinished := range finished {
        if !isFinished {
            unfinished = append(unfinished, todo_list[idx])
        }
    }

    return unfinished
}
