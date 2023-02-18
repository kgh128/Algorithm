package Baekjoon;

import java.io.*;

public class p9095 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        // 1. 테스트 케이스 수 입력받기
        int T = Integer.parseInt(br.readLine());

        // 2. 모든 n(1~11)에 대한 경우의 수 구하기
        // i = (i-3) + 3
        // i = (i-2) + 2
        // i = (i-1) + 1
        // 3칸씩 도는 규칙이므로 초기 3개의 값은 미리 넣어줌.
        int[] dp = new int[12];
        dp[1] = 1;
        dp[2] = 2;
        dp[3] = 4;

        for (int i = 4; i <= 11; i++) {
            dp[i] = dp[i-3] + dp[i-2] + dp[i-1];
        }

        // 3. 테스트 케이스 실행
        for (int i = 0; i < T; i++) {
            int n = Integer.parseInt(br.readLine());
            sb.append(dp[n]).append('\n');
        }

        // 4. 출력하기
        System.out.print(sb);
    }
}
