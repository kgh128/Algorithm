package Programmers;

import java.util.Arrays;

public class 체육복 {
    public int solution(int n, int[] losts, int[] reserves) {
        int answer = n;
        int[] student = new int[n+2];
        Arrays.fill(student, 1);

        // 1. 현재 상태 정리
        for (int lost: losts) {
            student[lost]--;
        }

        for (int reserve: reserves) {
            student[reserve]++;
        }

        // 2. 체육복 빌리기
        for (int i = 1; i <= n; i++) {
            if (student[i] == 0) {
                if (student[i-1] == 2) {
                    student[i]++;
                    student[i-1]--;
                }
                else if (student[i+1] == 2) {
                    student[i]++;
                    student[i+1]--;
                }
                else answer--;
            }
        }

        return answer;
    }
}
