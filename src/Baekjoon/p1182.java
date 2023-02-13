package Baekjoon;

import java.io.*;

public class p1182 {
    static int N;
    static int S;
    static int[] nums;
    static int count = 0;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        String[] inputs = br.readLine().split(" ");
        N = Integer.parseInt(inputs[0]);
        S = Integer.parseInt(inputs[1]);

        inputs = br.readLine().split(" ");
        nums = new int[N];

        for (int i = 0; i < N; i++) {
            nums[i] = Integer.parseInt(inputs[i]);
        }

        // 2. 부분수열의 합 구하기
        for (int i = 0; i < N; i++) {
            sumOfSubset(0, i);
        }

        // 3. 출력하기
        System.out.println(count);
    }

    public static void sumOfSubset(int sum, int current) {
        sum += nums[current];

        if (sum == S) {
            count++;
        }
        if (current == N-1) {
            return;
        }

        for (int i = current+1; i < N; i++) {
            sumOfSubset(sum, i);
        }
    }
}
