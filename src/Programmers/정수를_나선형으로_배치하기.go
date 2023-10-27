func solution(n int) [][]int {
    // 1. 시작 위치 => arr[0][-1]
    // 가장 처음 모드가 오른쪽으로 한 칸씩 가는 것이므로
    // arr[0][0]을 채워주기 위해서 시작 위치는 이론상 한 칸 왼쪽인 arr[0][-1]
    var x, y = -1, 0

    // 2. mode에 따른 진행 방향 설정
    // mode: 0, 1, 2, 3 => 우, 하, 좌, 상
    dx := []int{1, 0, -1, 0}
    dy := []int{0, 1, 0, -1}

    // 3. 배열에 넣을 값 초기화
    input := 1

    // 4. n*n 2차원 배열 생성
    arr := make([][]int, n)
    for i := 0; i < n; i++ {
        arr[i] = make([]int ,n)
    }

    // 5. 모드가 바뀌는 횟수: 2*n-1번
    // n = 1: 우 => 1번
    // n = 2: 우하좌 => 3번
    // n = 3: 우하좌상우 => 5번
    // n = 4: 우하좌상우히좌 => 7번
    // i를 0부터 시작한다고 하면 가장 마지막은 2*(n-1)
    for i := 0; i <= 2 * (n - 1); i++ {
        mode := i % 4            // 모드가 4개씩 순서대로 번갈아 가면서 수행됨.
        cnt := n - (i + 1) / 2   // 한 모드에서 값 채우기가 진행되는 횟수

        /*
        ex) n = 4
        <i>             =>     <i+1>
        mode 0: 0 4 8           1 5 9   
        mode 1: 1 5             2 6
        mode 2: 2 6             3 7
        mode 3: 3 7             4 8

        <cnt>
        mode 0: 5 3 1
        mode 1: 4 2
        mode 2: 4 2
        mode 3: 3 1 
        */

        for j := 0; j < cnt; j++ {
            // 6. 다음 위치로 이동
            x += dx[mode]
            y += dy[mode]

            // 7. 배열에 값 집어넣기
            arr[y][x] = input
            input++
        }
    }

    return arr
}
