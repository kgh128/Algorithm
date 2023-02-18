package Baekjoon;

import java.io.*;
import java.util.Arrays;

public class p1037 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int M = Integer.parseInt(br.readLine());
        int[] A = new int[M];
        String[] inputs = br.readLine().split(" ");

        for (int i = 0; i < M; i++) {
            A[i] = Integer.parseInt(inputs[i]);
        }

        // 2. 정렬하기
        Arrays.sort(A);

        // 3. N 구하기
        int N = A[0] * A[M-1];
        System.out.println(N);
    }
}
