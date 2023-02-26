package Programmers;

import java.util.PriorityQueue;

public class 더_맵게 {
    public int solution(int[] scoville, int K) {
        int answer = 0;
        PriorityQueue<Integer> heap = new PriorityQueue<>();

        // 1. 힙에 음식의 스코빌 지수 넣기
        for (int food: scoville) {
            heap.add(food);
        }

        // 2. 음식의 스코빌 지수를 K 이상으로 만들기
        // 힙에 음식이 2개 이상 들어있고, 최소 스코빌 지수가 K보다 작은 경우
        while (heap.size() > 1 && heap.peek() < K) {
            int firstMin = heap.poll();
            int secondMin = heap.poll();

            heap.add(firstMin + secondMin*2);
            answer++;
        }

        // 3. 결과 출력
        if (heap.poll() < K) return -1; // 최종 음식의 스코빌 지수가 K 미만이면 실패
        else return answer;             // 최종 음식의 스코빌 지수가 K 이상이면 성공
    }
}
