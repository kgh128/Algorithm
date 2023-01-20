package Baekjoon.Sorting;

import java.io.*;
import java.util.*;

public class p10815 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[] card = new int[N];
        String[] input = br.readLine().split(" ");

        for (int i = 0; i < N; i++) {
            card[i] = Integer.parseInt(input[i]);
        }

        int M = Integer.parseInt(br.readLine());
        int[] target = new int[M];
        input = br.readLine().split(" ");

        for (int i = 0; i < M; i++) {
            target[i] = Integer.parseInt(input[i]);
        }

        // 2. 카드를 오름차순으로 정렬하기
        Arrays.sort(card);

        // 3. 이분탐색
        for (int i = 0; i < M; i++) {
            if (Arrays.binarySearch(card, target[i]) >= 0) {
                bw.append("1 ");
            }
            else {
                bw.append("0 ");
            }
        }

        // 4. 결과 출력
        System.out.print(bw);
    }
}
