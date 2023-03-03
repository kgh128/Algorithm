package Baekjoon;

import java.io.*;
import java.util.*;

public class p2178 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        String[] line = br.readLine().split(" ");
        int N = Integer.parseInt(line[0]);
        int M = Integer.parseInt(line[1]);

        int[][] map = new int[N+2][M+2];    // 실제 지도를 0으로 감싸기
        int[] dx = {1, -1, 0, 0};           // 행 이동: 상하좌우
        int[] dy = {0, 0, -1, 1};           // 열 이동: 상하좌우

        for (int x = 1; x <= N; x++) {
            line = br.readLine().split("");

            for (int y = 1; y <= M; y++) {
                // 0: 길 X, -1: 길 O
                if (Integer.parseInt(line[y-1]) == 1) {
                    map[x][y] = -1;
                }
            }
        }

        // 2. bfs로 길 찾기
        Queue<Point> queue = new LinkedList<>();
        queue.add(new Point(1, 1));
        map[1][1] = 1; // 방문 표시 및 지나온 칸 수 저장

        while (!queue.isEmpty()) {
            Point current = queue.poll();

            // (N, M)에 도착
            if (current.x == N && current.y == M) {
                System.out.println(map[N][M]);
                break;
            }

            // 현재 칸의 상하좌우 탐색
            for (int i = 0; i < 4; i++) {
                int nx = current.x + dx[i];
                int ny = current.y + dy[i];

                if (map[nx][ny] == -1) {
                    queue.add(new Point(nx, ny));
                    map[nx][ny] = map[current.x][current.y] + 1; // 방문 표시 및 지나온 칸 수 저장
                }
            }
        }
    }

    public static class Point {
        int x;
        int y;

        public Point(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }
}
