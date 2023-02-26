package Baekjoon;

import java.io.*;
import java.util.*;

public class p1697 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        StringTokenizer st = new StringTokenizer(br.readLine());
        int N = Integer.parseInt(st.nextToken());
        int K = Integer.parseInt(st.nextToken());

        // 2. bfs
        Queue<Integer> queue = new LinkedList<>(); // 노드의 좌표값 저장
        int[] time = new int[100001]; // 중복 방문 방지 및 해당 노드에 도달하는데 걸리는 최소 시간 저장

        queue.add(N);
        Arrays.fill(time, -1);
        time[N] = 0;

        while (!queue.isEmpty()) {
            int node = queue.remove();

            // 동생을 찾은 경우
            if (node == K) {
                System.out.println(time[node]);
                break;
            }

            // X-1 이동 (X가 음수가 안되도록 함.)
            if (node > 0 && time[node-1] == -1) {
                time[node-1] = time[node] + 1;
                queue.add(node-1);
            }
            // X+1 이동
            if (node < 100000 && time[node+1] == -1) {
                time[node+1] = time[node] + 1;
                queue.add(node+1);
            }
            // 2*X 이동
            if (node <= 50000 && time[node*2] == -1) {
                time[node*2] = time[node] + 1;
                queue.add(node*2);
            }
        }
    }
}

/*
1. 중복 방문을 막지 않으면 방문했던 노드를 또 큐에 넣을 수 있음.
   그러면 같은 노드들이 너무 많이 큐에 들어가 메모리 초과 발생.

   ex) 9 + 1 = 10을 하여 노드 10을 큐에 넣음.
   ex) 5 * 2 = 10을 하여 노드 10을 큐에 넣음.
   위와 같이 10의 좌표값을 가지는 노드가 여러 번 큐에 들어갈 수 있음.

2. depth를 저장하는 큐와 중복 방문 방지 visited 배열을 time 배열로 합침. (다른 풀이보고 수정)
   time[node] == -1이면 아직 방문하지 않은 노드이니 방문 가능.
   time[node] > -1이면 이미 방문한 노드이고, time[node]의 값은 해당 node의 depth 값임.
   (그래프에서의 depth 값은 해당 노드에 도달하는데 걸리는 최소 시간)
*/