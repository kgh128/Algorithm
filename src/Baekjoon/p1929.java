package Baekjoon;

import java.io.*;

public class p1929 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        // 1. 입력받기 및 변수 선언
        String[] inputs = br.readLine().split(" ");
        int M = Integer.parseInt(inputs[0]);
        int N = Integer.parseInt(inputs[1]);
        boolean[] isNotPrime = new boolean[N+1];

        // 2. 소수 판별 및 배수 지우기
        for (int i = 2; i < N+1; i++) {
            if (!isNotPrime[i]) {
                if (i >= M) sb.append(i).append('\n');

                for (int j = 2; i*j < N+1; j++) {
                    isNotPrime[i*j] = true;
                }
            }
        }

        // 3. 출력하기
        System.out.print(sb);
    }
}
