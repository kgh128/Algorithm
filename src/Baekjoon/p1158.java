package Baekjoon;

import java.io.*;
import java.util.*;

public class p1158 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder("<");

        // 1. 입력받기
        String[] inputs = br.readLine().split(" ");
        int N = Integer.parseInt(inputs[0]);
        int K = Integer.parseInt(inputs[1]);

        // 2. 큐에 1~N 넣기
        Queue<Integer> queue = new LinkedList<>();

        for (int i = 1; i < N+1; i++) {
            queue.add(i);
        }

        // 3. 큐가 빌 때까지 순환하면서 K번째 원소 출력
        int count = 0;

        while (!queue.isEmpty()) {
            int head = queue.poll();
            count++;

            if (count == K) {
                count = 0;
                bw.append(head).append(", ");
            }
            else {
                queue.add(head);
            }
        }

        // 4. 결과 출력
        bw.replace(bw.length()-2, bw.length()-1, ">");
        System.out.print(bw);
    }
}
