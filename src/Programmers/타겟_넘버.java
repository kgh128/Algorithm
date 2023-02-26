package Programmers;

import java.util.*;

// 1) bfs 풀이
public class 타겟_넘버 {
    public int solution(int[] numbers, int target) {
        int answer = 0;
        int height = numbers.length - 1;
        Queue<Integer> valueQueue = new LinkedList<>();
        Queue<Integer> depthQueue = new LinkedList<>();

        // 초기값 넣어주기
        // 0 - numbers[0], 0 + numbers[0]도 할 수 있도록 0을 초기값으로 넣어주기(depth는 -1)
        valueQueue.add(0);
        depthQueue.add(-1);

        // bfs 수행
        while (!valueQueue.isEmpty()) {
            int nodeValue = valueQueue.poll();
            int nodeDepth = depthQueue.poll();

            // 연산이 끝나지 않았을 때
            if (nodeDepth < height) {
                // 1. 다음 수 빼기
                valueQueue.add(nodeValue - numbers[nodeDepth+1]);
                depthQueue.add(nodeDepth+1);

                // 2. 다음 수 더하기
                valueQueue.add(nodeValue + numbers[nodeDepth+1]);
                depthQueue.add(nodeDepth+1);
            }
            // 최종 결과가 타겟 넘버일 때
            else if (nodeValue == target) {
                answer++;
            }
        }

        return answer;
    }
}

// 2) dfs 풀이
/*
class 타겟_넘버 {
    public int solution(int[] numbers, int target) {
        // dfs 시행
        return dfs(numbers, target, -1, 0);
    }

    public int dfs(int[] numbers, int target, int depth, int value) {
        // 마지막 숫자까지 계산 완료했을 경우 - 종료 조건
        if (depth == numbers.length-1) {
            if (value == target) return 1;
            else return 0;
        }

        // 다음 수를 빼는 경우
        int minus = dfs(numbers, target, depth+1, value-numbers[depth+1]);

        // 다음 수를 더하는 경우
        int plus = dfs(numbers, target, depth+1, value+numbers[depth+1]);

        return minus + plus;
    }
}
*/