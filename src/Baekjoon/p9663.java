package Baekjoon;

import java.io.*;

// 1번 방법: for문       -> O(n)    => 2556ms
// 2번 방법: 이중 for문   -> O(n^2)  => 5636ms
// 실제 실행시간도 1번 방법이 2번 방법의 1/2임.

// 1. 퀸의 공격 경로를 표시하는 방법
public class p9663 {
    static int N, count;
    static boolean[] isAttacked1;   // 현재 좌표의 열이 퀸의 공격 경로인지
    static boolean[] isAttacked2;   // 현재 좌표를 지나는 기울기가 1인 대각선이 퀸의 공격 경로인지
    static boolean[] isAttacked3;   // 현재 좌표를 지나는 기울기가 -1인 대각선이 퀸의 공격 경로인지

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 변수 초기화
        N = Integer.parseInt(br.readLine());
        count = 0;
        isAttacked1 = new boolean[N];
        isAttacked2 = new boolean[2*N-1];
        isAttacked3 = new boolean[2*N-1];

        // 2. dfs로 완전탐색
        dfs(0);

        // 3. 출력하기
        System.out.println(count);
    }

    public static void dfs(int row) {
        if (row > N-1) {
            count++;
            return;
        }

        // 각 행의 열을 반복문으로 돌면서 확인
        for (int col = 0; col < N; col++) {
            // 현재 좌표가 이전 퀸들의 공격 경로에 있다면 놓을 수 없음.
            if (isAttacked1[col] || isAttacked2[row+col] || isAttacked3[row-col+N-1]) {
                continue;
            }

            // 현재 좌표가 이전 퀸들의 공격 경로에 있지 않다면 퀸을 놓을 수 있음.
            // 현재 좌표에 퀸을 놓았을 때 그 퀸의 공격 경로를 표시하고, 다음 행으로 넘어가기
            isAttacked1[col] = true;
            isAttacked2[row+col] = true;
            isAttacked3[row-col+N-1] = true;
            dfs(row+1);
            isAttacked1[col] = false;
            isAttacked2[row+col] = false;
            isAttacked3[row-col+N-1] = false;
        }
    }
}

/*
// 2. 퀸의 위치를 저장하고, 서로 공격 가능한 위치인지 확인하는 방법
public class p9663 {
    static int N, count;
    static int[] queen;   // 원소 인덱스: 행, 원소 값: 열 (각 행마다 어느 열에 퀸이 있는지를 저장)

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        N = Integer.parseInt(br.readLine());
        queen = new int[N]; // 한 행에는 한 개의 퀸만 있어야 하므로 배열의 크기는 행의 개수인 N이면 됨.

        // 2. dfs로 완전탐색
        count = 0;
        dfs(0);

        // 3. 출력하기
        System.out.println(count);
    }

    public static void dfs(int row) {
        if (row > N-1) {
            count++;
            return;
        }

        // 한 행에 대해 열마다 돌기
        for (int col = 0; col < N; col++) {
            // 퀸을 놓을 수 있으면 놓고 다음 행으로 넘어가기
            if (promising(row, col)) {
                queen[row] = col;
                dfs(row+1);
            }
        }
    }

    public static boolean promising(int row, int col) {
        // 0행이면 처음 놓는 퀸이므로 무조건 놓을 수 있음.
        if (row == 0) return true;

        // 0행이 아닌 경우 다른 퀸들과 같은 열 또는 대각선에 있는지 확인.
        for (int i = 0; i < row; i++) {
            if (queen[i] == col || Math.abs(row-i) == Math.abs(col-queen[i])) {
                return false;
            }
        }
        return true;
    }
}
*/