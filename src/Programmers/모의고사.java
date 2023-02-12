package Programmers;

import java.util.ArrayList;

public class 모의고사 {
    public int[] solution(int[] answers) {
        ArrayList<Integer> answer = new ArrayList<>();

        int[] answers1 = {1, 2, 3, 4, 5};
        int[] answers2 = {2, 1, 2, 3, 2, 4, 2, 5};
        int[] answers3 = {3, 3, 1, 1, 2, 2, 4, 4, 5, 5};

        int score1 = 0;
        int score2 = 0;
        int score3 = 0;

        for (int i = 0; i < answers.length; i++) {
            if (answers[i] == answers1[i%5]) score1++;
            if (answers[i] == answers2[i%8]) score2++;
            if (answers[i] == answers3[i%10]) score3++;
        }

        int maxScore = Math.max(score1, Math.max(score2, score3));

        if (score1 == maxScore) answer.add(1);
        if (score2 == maxScore) answer.add(2);
        if (score3 == maxScore) answer.add(3);

        return answer.stream().mapToInt(i->i).toArray();
    }
}
