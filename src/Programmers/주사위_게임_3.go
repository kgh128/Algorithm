import "sort"

func max(x int, y int) int {
    if (x > y) {
        return x
    } else {
        return y
    }
}

func abs(x int) int {
    if (x > 0) {
        return x
    } else {
        return 0 - x
    }
}

// 1. 정렬한 슬라이스 반복문 돌면서 같은 것의 개수 세기
func solution(a int, b int, c int, d int) int {
    var p, q, r, max_n = 0, 0, 0, 0
    dice := []int{a, b, c, d}

    sort.Ints(dice)

    for i := 0; i < 4; i++ {
        n := 1

        for j := i+1; j < 4; j++ {
            if (dice[i] == dice[j]) {
                n++  // 중복 숫자의 개수 세기
            } else {
                break
            }
        }

        // 1) 중복 숫자가 2개 이상 존재할 때 - p에 값이 저장되어야 함.
        if n > 1 {
            // p가 0일 때는 q에 넣지 않기 위해 추가
            // 어차피 슬라이스는 오름차순으로 정렬되어 있으므로
            // p가 0이 아니라면 무조건 p > q
            q = max(q, p)
            p = dice[i]
        } else {  // 2) 중복 숫자가 존재하지 않을 때 - p에 값이 저장될 필요X
            r = q
            q = dice[i]
        }

        max_n = max(max_n, n)  // 중복 숫자의 최대 개수 저장 -> 이후에 어떤 경우인지 판별할 때 사용
        i = i + n - 1          // 이미 j를 통해 돌아본 i들은 건너뛰기 (정렬되어 있으므로 j를 벗어난 이상, dice[j]와 다시 같은 숫자가 나올 일은 없음.)
    }

    switch max_n {
        case 4:
            return 1111 * p
        case 3:
            return (10 * p + q) * (10 * p + q)
        case 2:
            if r == 0 {
                return (p + q) * abs(p - q)
            } else {
                return q * r
            }
        default:
            return dice[0]
    }
}

// 2. 정렬한 슬라이스를 가지고 나올 수 있는 경우를 if문으로 나누기
func solution(a int, b int, c int, d int) int {
    dice := []int{a, b, c, d}
    sort.Ints(dice)

    // 1) 네 주사위에서 나온 숫자가 모두 같을 때
    // ex) 1 1 1 1
    if dice[0] == dice[3] {
        return 1111 * dice[0]

    // 2) 세 주사위에서 나온 숫자가 모두 같을 때 + 1개는 다름.
    // ex) 1 1 1 2 / 1 2 2 2
    } else if dice[0] == dice[2] {
        return (10 * dice[0] + dice[3]) * (10 * dice[0] + dice[3])
        
    } else if dice [1] == dice[3] {
        return (10 * dice[1] + dice[0]) * (10 * dice[1] + dice[0])

    // 3) 두 주사위에서 나온 숫자가 같고, 나머지 두 주사위에서 나온 숫자도 같을 때
    // ex) 1 1 2 2
    } else if dice[0] == dice[1] && dice[2] == dice[3] {
        return (dice[0] + dice[2]) * abs(dice[0] - dice[2])

    // 4) 두 주사위에서 나온 숫자만 같고, 나머지 두 주사위에서 나온 숫자는 서로 다를 때
    // ex) 1 1 2 3 / 1 2 2 3 / 1 2 3 3
    } else if dice[0] == dice[1] {
        return dice[2] * dice[3]
        
    } else if dice[1] == dice[2] {
        return dice[0] * dice[3]
        
    } else if dice[2] == dice[3] {
        return dice[0] * dice[1]

    // 4) 네 주사위에서 나온 숫자가 모두 다를 때
    // ex) 1 2 3 4
    } else {
        return dice[0]
    }
}
