package Baekjoon;

import java.io.*;

public class p1978 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[] nums = new int[N];
        String[] inputs = br.readLine().split(" ");

        for (int i = 0; i < N; i++) {
            nums[i] = Integer.parseInt(inputs[i]);
        }

        // 2. 1000이하의 소수 판별
        boolean[] isNotPrime = new boolean[1001];
        isNotPrime[1] = true;   // 1은 소수X

        for (int i = 2; i < 1001; i++) {
            if (!isNotPrime[i]) {
                for (int j = 2; i*j < 1001; j++) {
                    isNotPrime[i*j] = true;
                }
            }
        }

        // 3. 주어진 수들 중 소수의 개수 구하기
        int count = 0;

        for (int num: nums) {
            if (!isNotPrime[num]) count++;
        }

        // 4. 출력하기
        System.out.println(count);
    }
}
