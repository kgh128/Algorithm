// 두 풀이 모두 이전 배열과 현재 배열을 직접 비교하지 않았다.
// (정확히는 슬라이스)

/*
1. 변하지 않는 원소들의 개수 세기
- 조건에 충족되지 않아 연산을 수행하지 않는 원소들의 개수를 센다.
- 해당 개수가 전체 원소 개수와 같으면 배열의 아무 원소도 변하지 않았다는 의미이다.
=> 이때 반복문을 빠져나가서 x-1을 반환한다. 
    arr(x-1) = arr(x)이므로 x-1을 반환한다.
*/

func solution(arr []int) int {
    x := 0

    for notChange := 0; notChange < len(arr); x++ {
        notChange = 0

        for idx, num := range arr {
            if num >= 50 && num % 2 == 0 {
                arr[idx] = num / 2
            } else if num < 50 && num % 2 == 1 {
                arr[idx] = num * 2 + 1
            } else {
                notChange++
            }
        }
    }

    return x - 1
}


/*
2. 연산을 수행할 때 bool 변수에 배열이 변했음을 표시하기
- 조건이 충족되어 연산을 수행하면 원소의 값이 바뀌는 것이므로 이때 bool 변수에 배열이 변했음을 true로 표시한다.
- 만약 아무 연산도 수행되지 않으면 bool 변수가 false일 것이고, 이때는 배열이 이전 배열과 동일할 것이다.
- 그때 반복문을 빠져나가서 x-1을 반환한다.
*/

func solution(arr []int) int {
    x := 0
    
    for change := true; change; x++ {
        change = false
        
        for idx, num := range arr {
            if num >= 50 && num % 2 == 0 {
                arr[idx] = num / 2
                change = true
            } else if num < 50 && num % 2 == 1 {
                arr[idx] = num * 2 + 1
                change = true
            } 
        }
    }
    
    return x - 1
}
