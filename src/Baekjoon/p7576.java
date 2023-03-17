package Baekjoon;

import java.io.*;

public class p7576 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        String[] line = br.readLine().split(" ");
        int M = Integer.parseInt(line[0]);
        int N = Integer.parseInt(line[1]);
        int[][] storage = new int[N][M];

        for (int row = 0; row < N; row++) {
            line = br.readLine().split(" ");

            for (int col = 0; col < M; col++) {
                storage[row][col] = Integer.parseInt(line[col]);
            }
        }
    }
}
