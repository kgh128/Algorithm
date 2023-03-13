package Baekjoon;

import java.io.*;
import java.util.StringTokenizer;

public class p1149 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[][] dp = new int[N][3]; // 0: R, 1: G, 2: B

        for (int i = 0; i < N; i++) {
            StringTokenizer st = new StringTokenizer(br.readLine());

            dp[i][0] = Integer.parseInt(st.nextToken()); // R
            dp[i][1] = Integer.parseInt(st.nextToken()); // G
            dp[i][2] = Integer.parseInt(st.nextToken()); // B
        }

        // 2. 동적 프로그래밍 수행
        // 각 집을 R, G, B로 칠했을 때의 최솟값을 각각 구하기
        for (int i = 1; i < N; i++) {
            dp[i][0] += Math.min(dp[i-1][1], dp[i-1][2]); // R: G, B 중에 최소값 더하기
            dp[i][1] += Math.min(dp[i-1][0], dp[i-1][2]); // G: R, B 중에 최솟값 더하기
            dp[i][2] += Math.min(dp[i-1][0], dp[i-1][1]); // B: R, G 중에 최솟값 더하기
        }

        // 3. 비용의 최솟값 출력하기
        // 마지막 집을 R, G, B로 칠하는 것 중에서 최솟값을 구하기
        System.out.println(Math.min(dp[N-1][0], Math.min(dp[N-1][1], dp[N-1][2])));
    }
}
