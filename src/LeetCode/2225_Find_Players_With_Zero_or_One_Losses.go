import "sort"

func findWinners(matches [][]int) [][]int {
    lostMatches := make(map[int]int)
    answer := make([][]int, 2)

    for _, match := range matches {
        if _, exists := lostMatches[match[0]]; !exists {
            lostMatches[match[0]] = 0;
        }
        lostMatches[match[1]]++;
    }

    for player, count := range lostMatches {
        if count == 0 {
            answer[0] = append(answer[0], player);
        } else if count == 1 {
            answer[1] = append(answer[1], player);
        }
    }

    sort.Ints(answer[0]);
    sort.Ints(answer[1]);

    return answer;
}
