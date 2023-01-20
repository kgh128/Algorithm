package Baekjoon.Sorting;

import java.io.*;
import java.util.Arrays;

public class p10989 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int num[] = new int[N];

        for (int i = 0; i < N; i++) {
            num[i] = Integer.parseInt(br.readLine());
        }

        // 2. 오름차순으로 정렬하기
        Arrays.sort(num);

        // 3. 결과 출력하기
        for (int i = 0; i < N; i++) {
            bw.append(num[i]).append('\n');
        }
        System.out.print(bw);
    }
}
