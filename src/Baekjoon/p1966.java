package Baekjoon;

import java.io.*;
import java.util.*;

public class p1966 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 테스트 케이스 수행
        int testCase = Integer.parseInt(br.readLine());

        for (int i = 0; i < testCase; i++) {
            // 2. 입력받기
            String[] inputs = br.readLine().split(" ");
            int N = Integer.parseInt(inputs[0]);
            int M = Integer.parseInt(inputs[1]);

            inputs = br.readLine().split(" ");
            Integer[] priorities = new Integer[N];              // 정렬할 우선순위 배열
            Queue<List<Integer>> queue = new LinkedList<>();    // 연산에 사용할 큐 -> {번호, 중요도} 저장

            for (int j = 0; j < N; j++) {
                int priority = Integer.parseInt(inputs[j]);

                priorities[j] = priority;
                queue.add(Arrays.asList(j, priority));
            }

            // 3. 중요도 내림차순 정렬하기
            Arrays.sort(priorities, Comparator.reverseOrder());

            // 4. 큐 연산하기
            int count = 0;

            while (true) {
                List<Integer> head = queue.remove();

                if (head.get(1) < priorities[count]) {
                    queue.add(head);
                }
                else {
                    count++;

                    if (head.get(0) == M) {
                        bw.append(count).append('\n');
                        break;
                    }
                }
            }
        }
        System.out.print(bw);
    }
}
