// 1. 선형 탐색 - O(N)의 시간복잡도: 15ms
func mySqrt(x int) int {
    if x == 0 || x == 1 {
        return x
    }

    for i := 2; i <= x; i++ {
        square := i*i

        if x == square {
            return i
        } else if x < square {
            return i-1
        }
    }

    return 0
}

// 2. 이진 탐색 - O(logN)의 시간복잡도: 0ms
func mySqrt(x int) int {
    left, right := 0, x

    for left <= right {
        mid := (left + right) / 2

        if x < mid * mid {
            right = mid - 1
        } else if x > mid * mid {
            left = mid + 1
        } else {
            return mid
        }
    }

    return right
}

// 3. math 라이브러리 사용 - 0ms
import "math"

func mySqrt(x int) int {
    return int(math.Sqrt(float64(x)))
}
