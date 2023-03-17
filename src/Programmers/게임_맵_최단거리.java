package Programmers;

import java.util.Queue;
import java.util.LinkedList;

public class 게임_맵_최단거리 {
    class Point {
        int x;
        int y;

        Point(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }

    public int solution(int[][] maps) {
        int n = maps.length;
        int m = maps[0].length;
        int[] dx = {0, 0, -1, 1}; // 동서남북 행 이동 방향
        int[] dy = {1, -1, 0, 0}; // 동서남북 열 이동 방향
        Queue<Point> queue = new LinkedList<>();

        // bfs 수행
        queue.add(new Point(0, 0));

        while(!queue.isEmpty()) {
            Point current = queue.poll();

            for (int i = 0; i < 4; i++) {
                int nx = current.x + dx[i]; // 다음 노드의 행
                int ny = current.y + dy[i]; // 다음 노드의 열

                if (nx >= 0 && nx < n && ny >= 0 && ny < m) {
                    if (!(nx == 0 && ny == 0) && maps[nx][ny] == 1) {
                        maps[nx][ny] = maps[current.x][current.y] + 1; // 최단 거리 저장 및 방문 표시
                        queue.add(new Point(nx, ny));
                    }
                }

                // 상대 팀 진영에 도착한 경우
                if (nx == n-1 && ny == m-1) {
                    return maps[nx][ny];
                }
            }
        }

        // 상대 팀 진영에 도착할 수 없는 경우
        return -1;
    }
}
