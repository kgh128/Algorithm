// 1. integer를 string으로 바꾼 후 양끝값부터 중앙에 있는 값까지 서로 대칭인지 확인
import "strconv"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    strX := strconv.Itoa(x)

    for i := 0; i < len(strX)/2; i++ {
        if strX[i] != strX[len(strX)-i-1] {
            return false
        }
    }
    return true
}

// 2. 숫자 계산으로만 뒤집은 절반을 만들어서 기존 절반과 서로 같은지 비교
// ex) 12321 => reverseX: 123, x: 12 => x == reverseX / 10
// ex) 1221 => reverseX: 12, x: 12 => x == reverseX
func isPalindrome(x int) bool {
    // x가 음수이면 마이너스 기호 때문에 대칭이 아님.
    // x가 0으로 끝나는데 대칭이려면 0으로 시작해야 함. -> ex) 012210
    // 그러나 0으로 시작하는 수는 존재하지 않으므로 false 반환. 
    // (단, x가 0인 경우는 가능하므로 이 경우는 제외)
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }

    // x의 절반 중 뒤집은 수를 저장할 변수
    // x의 일의 자리수부터 절반까지 순서대로 더해준다.
    reverseX := 0

    
    for x > reverseX {
        // <뒤에서부터 뒤집어서 절반을 만들고, 앞에서부터 절반을 만들기>
        // 현재 reverseX에 10을 곱해서 일의 자리에 공간을 하나 만들어주고, (왼쪽으로 한 칸 밀어주기)
        // 그 일의 자리에 현재 x의 일의 자릿수를 더한다.
        // 그리고 x의 일의 자릿수는 더이상 필요 없으니 잘라내기 (오른쪽으로 한 칸 밀어주기)
        // ex) x = 12321
        // 1) reverseX = 0 * 10 + 1 = 1
        //    x = 1232
        // 2) reverseX = 1 * 10 + 2 = 12
        //    x = 123
        // 3) reverseX = 12 * 10 + 3 = 123
        //    x = 12
        // x < reverseX이므로 반복문 탈출 (이때는 x가 홀수 자리의 대칭인 수)
        // x가 짝수 자리의 대칭인 수이면 x == reverseX가 되어 반복문을 탈출할 것
        // x가 대칭인 수가 아니면 reverseX가 x보다 자리수가 많아질 때까지만 반복문을 돌고 탈출할 것
        reverseX = reverseX * 10 + x % 10
        x /= 10
    }

    // x가 짝수 자리의 대칭인 수이면 완전히 똑같은 절반으로 자를 수 있으므로 x == reverseX
    // x가 홀수 자리인 대칭인 수이면 reverseX의 일의 자릿수를 잘라낸 수와 x가 동일할 것이므로 x == reverseX / 10
    // (reverseX가 완전히 똑같은 절반에 가장 중앙에 있는 값 하나를 일의 자릿수로 가지고 있을 것이다.)
    return x == reverseX || x == reverseX / 10  
}
