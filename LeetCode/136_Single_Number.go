/*
<xor 연산 이용>
- xor 연산: 서로 다른 비트(0과 1)일 때 결과값으로 1을 반환
- 같은 숫자를 두 번 xor하면 원래 숫자로 돌아간다. (해당 숫자의 영향은 사라진다.)
- 배열의 숫자들을 순차적으로 xor하면 두 번 존재하는 숫자들은 두 번 xor 되어서 사라지고, 한 번만 존재하는 숫자만 남게된다.
*/

func singleNumber(nums []int) int {
    mask := nums[0];

    for i := 1; i < len(nums); i++ {
        mask ^= nums[i]
    }

    return mask
}
