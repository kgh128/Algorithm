package Baekjoon;

import java.io.*;

public class p14889 {
    static int N;
    static int[][] S;
    static boolean[] isStartTeam;
    static int minGap = 40000;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        N = Integer.parseInt(br.readLine());
        S = new int[N][N];
        isStartTeam = new boolean[N];

        for (int i = 0; i < N; i++) {
            String[] inputs = br.readLine().split(" ");

            for (int j = 0; j < N; j++) {
                S[i][j] = Integer.parseInt(inputs[j]);
            }
        }

        // 2. 최소 능력치 차이 얻기
        getMinGap(0, 0);

        // 3. 출력하기
        System.out.println(minGap);
    }

    public static void getMinGap(int current, int count) {
        if (count == N/2) {
            int startTeamSum = 0;
            int linkTeamSum = 0;

            // 팀의 종합 능력치 계산
            for (int i = 0; i < N; i++) {
                for (int j = i+1; j < N; j++) {
                    if (isStartTeam[i] && isStartTeam[j]) {
                        startTeamSum += S[i][j];
                        startTeamSum += S[j][i];
                    }
                    else if (!isStartTeam[i] && !isStartTeam[j]) {
                        linkTeamSum += S[i][j];
                        linkTeamSum += S[j][i];
                    }
                }
            }

            // 최소 능력치 차이 계산
            minGap = Math.min(minGap, Math.abs(startTeamSum - linkTeamSum));
            return;
        }

        // start 팀의 다음 선수 뽑기
        for (int i = current; i < N; i++) {
            isStartTeam[i] = true;
            getMinGap(i+1, count+1);
            isStartTeam[i] = false;
        }
    }
}