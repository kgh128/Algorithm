func plusOne(digits []int) []int {
    tens := 1

    for i := len(digits) - 1; i >= 0; i-- {
        sum := digits[i] + tens
        tens = sum / 10
        digits[i] = sum % 10

        if tens == 0 {
            return digits
        }
    }

    if tens == 1 {
        return append([]int{1}, digits[:]...)
    } else {
        return digits
    }
}
