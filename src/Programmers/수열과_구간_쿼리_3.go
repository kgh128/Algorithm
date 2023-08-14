/*
<Go에서의 swap>
a, b = b, a
위와 같이 써주기만 하면 a랑 b가 swap된다.
*/

func solution(arr []int, queries [][]int) []int {
    for _, query := range queries {
        i := query[0]
        j := query[1]

        arr[i], arr[j] = arr[j], arr[i]
    }

    return arr
}
