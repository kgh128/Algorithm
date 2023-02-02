package Baekjoon;

import java.io.*;
import java.util.Arrays;

public class p10867 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[] num = new int[N];
        String[] inputs = br.readLine().split(" ");

        for (int i = 0; i < N; i++) {
            num[i] = Integer.parseInt(inputs[i]);
        }

        // 2. 정렬하기
        Arrays.sort(num);

        // 3. 중복 제거하고 출력하기
        bw.append(num[0]).append(' ');

        for (int i = 1; i < N; i++) {
            if (num[i] == num[i-1]) {
                continue;
            }
            bw.append(num[i]).append(' ');
        }

        System.out.print(bw);
    }
}
