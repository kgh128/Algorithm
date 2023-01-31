package Baekjoon;

import java.io.*;
import java.util.*;

public class p11651 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[][] pairs = new int[N][2];

        for (int i = 0; i < N; i++) {
            String[] inputs = br.readLine().split(" ");

            pairs[i][0] = Integer.parseInt(inputs[0]);
            pairs[i][1] = Integer.parseInt(inputs[1]);
        }

        // 2. 정렬하기
        Arrays.sort(pairs, (a ,b) -> {
            if (a[1] == b[1]) return a[0] - b[0];
            else return a[1] - b[1];
        });

        // 3. 출력하기
        for (int i = 0; i < N; i++) {
            bw.append(pairs[i][0]).append(' ').append(pairs[i][1]).append('\n');
        }
        System.out.print(bw);
        //다시
    }
}
