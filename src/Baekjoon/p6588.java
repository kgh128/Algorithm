package Baekjoon;

import java.io.*;
import java.util.ArrayList;

public class p6588 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        // 1. 소수 판별 - 에라토스테네스의 체
        boolean[] isNotPrime = new boolean[1000001];
        ArrayList<Integer> oddPrimes = new ArrayList<>();

        isNotPrime[0] = true;
        isNotPrime[1] = true;

        for (int i = 2; i < 1000001; i++) {
            if (!isNotPrime[i]) {
                for (int j = 2; i*j < 1000001; j++) {
                    isNotPrime[i*j] = true;
                }

                if (i > 2) oddPrimes.add(i);
            }
        }

        // 2. 테스트 케이스 실행
        int n = Integer.parseInt(br.readLine());

        while (n > 0) {
            boolean possible = false;

            for (int oddPrime: oddPrimes) {
                int rest = n - oddPrime;

                if (!isNotPrime[rest]) {
                    sb.append(n).append(" = ").append(oddPrime).append(" + ").append(rest).append("\n");
                    possible = true;
                    break;
                }
            }

            if (!possible) {
                sb.append("Goldbach's conjecture is wrong.\n");
            }

            n = Integer.parseInt(br.readLine());
        }

        // 3. 출력하기
        System.out.print(sb);
    }
}
