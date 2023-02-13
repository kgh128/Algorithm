package Baekjoon;

import java.io.*;

public class p6603 {
    static BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
    static StringBuilder sb = new StringBuilder();
    static int[] S;
    static int K;

    public static void main(String[] args) throws IOException {
        // 1. 테스트 케이스 실행
        String[] inputs = br.readLine().split(" ");
        K = Integer.parseInt(inputs[0]);
        S = new int[K];

        while (K != 0) {
            // 3. 집합 S 저장하기
            for (int i = 0; i < K; i++) {
                S[i] = Integer.parseInt(inputs[i+1]);
            }

            // 4. 조합 만들기
            for (int i = 0; i < K; i++) {
                makeCombination(i, 0, new int[6]);
            }
            sb.append('\n');

            // 5. 다음 테스트 케이스 입력받기
            inputs = br.readLine().split(" ");
            K = Integer.parseInt(inputs[0]);
            S = new int[K];
        }

        // 6. 출력하기
        System.out.print(sb);
    }

    public static void makeCombination(int current, int count, int[] combination) {
        combination[count++] = S[current];

        if (count > 5) {
            for (int i = 0; i < 6; i++) {
                sb.append(combination[i]).append(' ');
            }
            sb.append('\n');

            return;
        }

        for (int i = current+1; i < K; i++) {
            makeCombination(i, count, combination);
        }
    }
}
