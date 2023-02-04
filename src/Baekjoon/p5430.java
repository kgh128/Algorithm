package Baekjoon;

import java.io.*;
import java.util.*;

public class p5430 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 테스트 케이스 실행
        int T = Integer.parseInt(br.readLine());

        for (int i = 0; i < T; i++) {
            // 2. 입력 & 파싱
            String[] functions = br.readLine().split("");
            int arraySize = Integer.parseInt(br.readLine());

            String inputs = br.readLine();
            String[] array = inputs.substring(1, inputs.length()-1).split(",");

            // 3. 덱에 배열 입력
            Deque<Integer> deque = new ArrayDeque<>(arraySize);

            if (arraySize > 0) {
                for (String num: array) {
                    deque.add(Integer.parseInt(num));
                }
            }

            // 4. 함수 수행
            boolean isReverseDeque = false;

            try {
                for (String func : functions) {
                    if (func.equals("R")) {
                        isReverseDeque = !isReverseDeque;
                    }
                    else if (func.equals("D") && !isReverseDeque) {
                        deque.removeFirst();
                    }
                    else {
                        deque.removeLast();
                    }
                }
            } catch (Exception e) {
                bw.append("error\n");
                continue;
            }

            // 5. 최종 결과 배열 저장
            bw.append("[");
            int dequeSize = deque.size();

            if (!isReverseDeque) {
                for (int j = 0; j < dequeSize; j++) {
                    bw.append(deque.removeFirst()).append(',');
                }
            }
            else {
                for (int j = 0; j < dequeSize; j++) {
                    bw.append(deque.removeLast()).append(',');
                }
            }

            if (dequeSize > 0) {
                // 마지막 쉼표 제거 (deque가 비어있는 경우는 쉼표가 아예 없어서 건너뜀.)
                bw.deleteCharAt(bw.length() - 1);
            }
            bw.append("]\n");
        }

        // 6. 결과 출력
        System.out.print(bw);
    }
}
