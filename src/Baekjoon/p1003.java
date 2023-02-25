package Baekjoon;

import java.io.*;

public class p1003 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        // 1. 테스트 케이스 실행
        int T = Integer.parseInt(br.readLine());
        int[] dp0 = new int[41];    // 0이 출력되는 횟수
        int[] dp1 = new int[41];    // 1이 출력되는 횟수

        // 2. 0번째 피보나치 수 결과 저장
        dp0[0] = 1;
        dp1[0] = 0;

        // 3. 1번째 피보나치 수 결과 저장
        dp0[1] = 0;
        dp1[1] = 1;

        for (int i = 0; i < T; i++) {
            // 4. 입력받기
            int N = Integer.parseInt(br.readLine());

            // 5. 동적 프로그래밍 수행
            for (int j = 2; j <= N; j++) {
                dp0[j] = dp0[j-1] + dp0[j-2];
                dp1[j] = dp1[j-1] + dp1[j-2];
            }

            // 6. 결과 저장
            sb.append(dp0[N]).append(" ").append(dp1[N]).append("\n");
        }

        // 7. 출력하기
        System.out.print(sb);
    }
}
