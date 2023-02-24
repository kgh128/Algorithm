package Baekjoon;

import java.io.*;
import java.util.ArrayList;

public class p1644 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력 받기
        int N = Integer.parseInt(br.readLine());

        // 2. 소수 배열 만들기 - 에라토스테네스의 체
        boolean[] isNotPrime = new boolean[N+1];
        ArrayList<Integer> primes = new ArrayList<>();

        isNotPrime[0] = true;
        isNotPrime[1] = true;

        for (int i = 2; i < N+1; i++) {
            if (!isNotPrime[i]) {
                primes.add(i);

                for (int j = 2; i*j < N+1; j++) {
                    isNotPrime[i*j] = true;
                }
            }
        }

        // 3. 연속된 소수의 합으로 나타낼 수 있는 경우의 수 구하기
        // 포인터를 2개 사용하여 소수 배열 탐색하기 - front, back
        int front = 0;
        int back = 0;
        int count = 0;
        int sum = 2;

        // N이 1인 경우는 sum > N이 항상 성립하므로 런타임 에러가 발생한다.
        // 따라서 N > 1인 경우만 while문 수행
        while (N > 1 && front <= back) {
            if (sum <= N) {
                if (sum == N) count++;

                if (back < primes.size()-1) {
                    sum += primes.get(++back);
                }
                else break; // 종료 조건 - sum <= N인데, back이 소수 배열 끝까지 왔으면 더이상 더할 소수가 없음.
            }
            else {
                sum -= primes.get(front++);
            }
        }

        // 4. 출력하기
        System.out.println(count);
    }
}