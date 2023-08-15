func solution(names []string) []string {
    firsts := []string{}

    for i := 0; i < len(names); i += 5 {
        firsts = append(firsts, names[i])
    }

    return firsts
}
