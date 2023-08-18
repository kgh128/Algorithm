func solution(ranks []int, attendance []bool) int {
    students := make([]int, len(ranks))
    selection := make([]int, 0, 3)

    // 1. students 배열에 rank가 높은 순으로 집어넣기
    // ranks 배열은 student-rank로 저장되어 있는데, 이를 rank-student로 뒤집어서 저장
    // => 인덱스: rank, 값: student
    // => 그러면 students 배열을 인덱스 0부터 순회하면 rank가 높은 순으로 student를 순회할 수 있음.
    for student, rank := range ranks {
        students[rank-1] = student
    }

    // 2. rank가 높은 순으로 순회하면서 참석한 사람 중 rank가 제일 높은 3명 뽑기
    for rank := 0; len(selection) < 3; rank++ {
        if attendance[students[rank]] {
            selection = append(selection, students[rank])
        }
    }

    // 3. 리턴값 계산
    return 10000 * selection[0] + 100 * selection[1] + selection[2]
}
