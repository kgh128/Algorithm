package Baekjoon;

import java.io.*;
import java.util.*;

public class p2667 {
    static int N;
    static int[][] map;
    static int count;
    static ArrayList<Integer> house;
    static int[] dx = {-1, 1, 0, 0};    // 행 이동 방향 - 상하좌우
    static int[] dy = {0, 0, -1, 1};    // 열 이동 방향 - 상하좌우

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        // 1. 입력받기
        N = Integer.parseInt(br.readLine());
        map = new int[N+2][N+2];
        house = new ArrayList<>();

        for (int i = 1; i <= N; i++) {
            String[] inputs = br.readLine().split("");

            for (int j = 1; j <= N; j++) {
                map[i][j] = Integer.parseInt(inputs[j-1]);
            }
        }

        // 2. dfs
        for (int i = 1; i <= N; i++) {
            for (int j = 1; j <= N; j++) {
                // 한 단지 내의 집의 수 세기
                if (map[i][j] == 1) {
                    count = 1;
                    map[i][j] = 0;  // 방문함.
                    dfs(i, j);
                    house.add(count);
                }
            }
        }

        // 3. 단지 내 집의 수를 오름차순 정렬
        house.sort(Comparator.naturalOrder());

        // 4. 출력하기
        sb.append(house.size()).append("\n");

        for (int houseCount: house) {
            sb.append(houseCount).append("\n");
        }

        System.out.print(sb);
    }

    public static void dfs(int x, int y) {
        // 상하좌우 확인
        for (int i = 0; i < 4; i++) {
            int nx = x + dx[i];
            int ny = y + dy[i];

            if (map[nx][ny] == 1) {
                count++;
                map[nx][ny] = 0;    // 방문함.
                dfs(nx, ny);
            }
        }
    }
}