/*
<for-range문>
1. 배열의 인덱스와 원소를 모두 사용
for idx, val := range array {...}

2. 배열의 인덱스만 사용
for idx := range array {...}

3. 배열의 원소만 사용
for _, val := range array {...}
*/

func solution(arr []int, k int) []int {
    if k%2 == 1 {
        for i := range arr {
            arr[i] *= k 
        }
    } else {
        for i := range arr {
            arr[i] += k
        }
    }
    
    return arr
}
