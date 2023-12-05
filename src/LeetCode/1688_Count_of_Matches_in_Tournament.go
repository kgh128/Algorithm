func numberOfMatches(n int) int {
    if n < 2 {
        return 0
    }

    curr := n / 2
    next := 0

    if n % 2 == 0 {
        next = numberOfMatches(curr)
    } else {
        next = numberOfMatches(curr + 1)
    }

    return curr + next
}
