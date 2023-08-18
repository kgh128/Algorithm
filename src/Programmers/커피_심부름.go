import "strings"

func solution(orders []string) int {
    cost := 0

    for _, order := range orders {
        if strings.Contains(order, "cafelatte") {
            cost += 5000
        } else {
            cost += 4500
        }
    }

    return cost
}
