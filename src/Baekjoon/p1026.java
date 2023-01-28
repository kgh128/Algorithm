package Baekjoon;

import java.io.*;
import java.util.*;
import java.lang.*;

public class p1026 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int S = 0;
        int N = Integer.parseInt(br.readLine());
        Integer[] A = new Integer[N];
        Integer[] B = new Integer[N];

        String[] inputA = br.readLine().split(" ");
        String[] inputB = br.readLine().split(" ");

        for (int i = 0; i < N; i++) {
            A[i] = Integer.parseInt(inputA[i]);
            B[i] = Integer.parseInt(inputB[i]);
        }

        // 2. 정렬
        Arrays.sort(A);                             // 오름차순 정렬
        Arrays.sort(B, Collections.reverseOrder()); // 내림차순 정렬

        // 3. S의 최솟값 계산
        for (int i = 0; i < N; i++) {
            S += A[i] * B[i];
        }

        System.out.println(S);
    }
}
