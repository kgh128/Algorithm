package Baekjoon;

import java.io.*;

public class p1932 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int n = Integer.parseInt(br.readLine());
        int[][] dp = new int[n][n];

        for (int i = 0; i < n; i++) {
            String[] line = br.readLine().split(" ");

            for (int j = 0; j < i+1; j++) {
                dp[i][j] = Integer.parseInt(line[j]);
            }
        }

        // 2. 동적 프로그래밍 - 현재 위치에 올 수 있는 합의 최댓값을 구함.
        for (int i = 1; i < n; i++) {
            for (int j = 0; j < i+1; j++) {
                // 가장 왼쪽에 있는 수(대각선 오른쪽만 접근 가능)
                if (j == 0) dp[i][j] += dp[i-1][j];

                // 가장 오른쪽에 있는 수(대각선 왼쪽만 접근 가능)
                else if (j == i) dp[i][j] += dp[i-1][j-1];

                // 중간에 있는 수(대각선 왼쪽과 대각선 오른쪽 모두 접근 가능)
                else dp[i][j] += Math.max(dp[i-1][j-1], dp[i-1][j]);
            }
        }

        // 3. 가장 아래층에 있는 dp값 중에서 최댓값 구하기
        int maxSum = 0;

        for (int j = 0; j < n; j++) {
            maxSum = Math.max(maxSum, dp[n-1][j]);
        }

        // 4. 출력하기
        System.out.println(maxSum);
    }
}