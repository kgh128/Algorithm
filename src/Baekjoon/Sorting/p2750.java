package Baekjoon.Sorting;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class p2750 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[] num = new int[N];

        for (int i = 0; i < N; i++) {
            num[i] = Integer.parseInt(br.readLine());
        }

        // 2. 오름차순 정렬
        Arrays.sort(num);

        // 3. 결과 출력
        for (int i = 0; i < N; i++) {
            bw.append(num[i]).append('\n');
        }
        System.out.print(bw);
    }
}
