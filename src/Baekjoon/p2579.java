package Baekjoon;

import java.io.*;

public class p2579 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[][] dp = new int[N][2];

        for (int i = 0; i < N; i++) {
            int score = Integer.parseInt(br.readLine());

            dp[i][0] = score;
            dp[i][1] = score;
        }

        // 2. 동적 프로그래밍 수행
        // dp[i][0]: i-1칸 O, i-2칸 X 경우의 점수 최댓값
        // dp[i][1]: i-1칸 X, i-2칸 O 경우의 점수 최댓값
        if (N > 1) dp[1][0] += dp[0][1]; // 계단이 1개 있는 경우는 할 필요X

        for (int i = 2; i < N; i++) {
            dp[i][0] += dp[i-1][1];
            dp[i][1] += Math.max(dp[i-2][0], dp[i-2][1]);
        }

        // 3. 총 점수의 최댓값 출력하기
        System.out.println(Math.max(dp[N-1][0], dp[N-1][1]));
    }
}
