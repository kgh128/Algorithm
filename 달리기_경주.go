func solution(players []string, callings []string) []string {
    find := make(map[string]int)

    for rank, name := range players {
        find[name] = rank
    }

    for _, name := range callings {
        rank := find[name]
        players[rank], players[rank-1] = players[rank-1], players[rank]
        find[name] = rank - 1
        find[players[rank]] = rank
    }

    return players
}
