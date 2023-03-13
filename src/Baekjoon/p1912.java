package Baekjoon;

import java.io.*;
import java.util.StringTokenizer;

public class p1912 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int n = Integer.parseInt(br.readLine());
        int[] dp = new int[n];
        StringTokenizer st = new StringTokenizer(br.readLine());

        for (int i = 0; i < n; i++) {
            dp[i] = Integer.parseInt(st.nextToken());
        }

        // 2. 동적 프로그래밍 수행
        for (int i = 1; i < n; i++) {
            dp[i] = Math.max(dp[i], dp[i-1] + dp[i]);
        }

        // 3. 연속합의 최댓값 구하기
        // dp 값 중에서 최댓값을 구하기
        int maxSum = dp[0];

        for (int i = 1; i < n; i++) {
            maxSum = Math.max(maxSum, dp[i]);
        }

        // 4. 출력하기
        System.out.println(maxSum);
    }
}
