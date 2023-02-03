package Baekjoon;

import java.io.*;
import java.util.*;

public class p1874 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        Queue<Integer> series = new LinkedList<>(); // 입력 수열 저장

        for (int i = 0; i < N; i++) {
            series.add(Integer.parseInt(br.readLine()));
        }

        // 2. 스택 연산 수행
        Stack<Integer> stack = new Stack<>();   // 연산 수행용 스택
        int head = series.remove();             // series(수열)에서 가장 앞에 있는 수
        int num = 1;                            // 스택에 넣을 수
        boolean popHeadAtStack = false;         // series에서 뽑은 head를 스택에서 제거했는지

        while(true) {
            // 수열을 다 만든 경우 (마지막 head까지 스택에서 제거)
            if (series.isEmpty() && popHeadAtStack) {
                break;
            }
            // 이전 head를 스택에서 제거해서 다음 head를 뽑아야하는 경우
            else if (!series.isEmpty() && popHeadAtStack) {
                head = series.remove();
            }

            // 스택에 숫자를 넣어야 하는 경우
            // 지금 스택에 넣어야 하는 수가 원하는 숫자(head)보다 작거나 같다 -> 아직 원하는 숫자가 스택에 안들어갔음.
            if (num <= head) {
                stack.push(num++);
                bw.append("+\n");
                popHeadAtStack = false;
            }
            // 스택에서 숫자를 빼야 하는 경우
            // 원하는 숫자(head)가 스택 가장 위에 있어서 뽑을 수 있음.
            else if (stack.peek() == head) {
                stack.pop();
                bw.append("-\n");
                popHeadAtStack = true;
            }
            // 입력 수열을 만들 수 없는 경우
            // 원하는 숫자(head)가 스택에 들어가 있기는 하지만, 가장 위에 있지는 않아서 뽑을 수 없음.
            else if (stack.peek() > head) {
                bw = new StringBuilder("NO");
                break;
            }
        }
        System.out.print(bw);
    }
}
