package Baekjoon;

import java.io.*;
import java.util.Stack;

public class p9012 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 테스트 케이스 실행하기
        int T = Integer.parseInt(br.readLine());

        for (int i = 0; i < T; i++) {
            // 2. 입력받기
            String ps = br.readLine();

            // 3. 스택 연산
            Stack<Character> stack = new Stack<>();
            boolean isVps = true;

            for (char p: ps.toCharArray()) {
                if (p == '(') {
                    stack.push(p);
                }
                else if (p == ')' && !stack.empty()) {
                    stack.pop();
                }
                else if (p == ')' && stack.empty()) {
                    isVps = false;
                    break;
                }
            }

            // 4. 스택 연산을 모두 수행했는데, 스택에 '('가 남은 경우 (짝을 못찾은 경우)
            if (!stack.empty()) {
                isVps = false;
            }

            // 4. 결과 출력
            if (isVps) bw.append("YES\n");
            else bw.append("NO\n");
        }
        System.out.print(bw);
    }
}
