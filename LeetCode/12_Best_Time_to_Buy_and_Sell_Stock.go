func maxProfit(prices []int) int {
    minPrice := prices[0]  // 최소 가격
    maxProfit := 0  // 최대 이익

    for i := 0; i < len(prices); i++ {
        if prices[i] < minPrice {
          // 현재 값이 최솟값보다 작은 경우 갱신
            minPrice = prices[i]
        } else if prices[i] - minPrice > maxProfit {
          // 현재 얻을 수 있는 이익이 최대 이익보다 큰 경우 갱신
            maxProfit = prices[i] - minPrice
        }
    }
    
    return maxProfit
}
