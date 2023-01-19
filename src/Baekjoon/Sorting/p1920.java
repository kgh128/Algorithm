package Baekjoon.Sorting;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class p1920 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        long[] stNum = new long[N];   // standard numbers
        String[] buffer = br.readLine().split(" ");

        for (int i = 0; i < N; i++) {
            stNum[i] = Long.parseLong(buffer[i]);
        }

        int M = Integer.parseInt(br.readLine());
        long[] cpNum = new long[M];   // compared numbers
        buffer = br.readLine().split(" ");

        for (int i = 0; i < M; i++) {
            cpNum[i] = Long.parseLong(buffer[i]);
        }

        // 2. stNum 배열 오름차순 정렬 - 이분탐색을 쓰기 위해서
        Arrays.sort(stNum);

        // 3. 이분탐색으로 stNum에 cpNum이 있는지 찾기
        for (int i = 0; i < M; i++) {
            if (Arrays.binarySearch(stNum, cpNum[i]) >= 0) {
                bw.append("1\n");
            }
            else {
                bw.append("0\n");
            }
        }

        // 4. 답 출력
        System.out.print(bw);
    }
}
