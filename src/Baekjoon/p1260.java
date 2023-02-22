package Baekjoon;

import java.io.*;
import java.util.*;

public class p1260 {
    static ArrayList<Integer>[] adj;
    static boolean[] visited;
    static Queue<Integer> queue;
    static StringBuilder sb;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        sb = new StringBuilder();

        // 1. 입력받기
        StringTokenizer st = new StringTokenizer(br.readLine());
        int N = Integer.parseInt(st.nextToken(" "));
        int M = Integer.parseInt(st.nextToken(" "));
        int V = Integer.parseInt(st.nextToken(" "));

        // 2. 그래프 만들기
        adj = new ArrayList[N+1];

        for (int i = 0; i < N+1; i++) {
            adj[i] = new ArrayList<>();
        }

        for (int i = 0; i < M; i++) {
            st = new StringTokenizer(br.readLine());
            int v1 = Integer.parseInt(st.nextToken(" "));
            int v2 = Integer.parseInt(st.nextToken(" "));

            adj[v1].add(v2);
            adj[v2].add(v1);
        }

        // 3. 정점 번호가 작은 인접 정점부터 방문해야 하므로 오름차순 정렬
        for (int i = 1; i < N+1; i++) {
            adj[i].sort(Comparator.naturalOrder());
        }

        // 4. dfs
        visited = new boolean[N+1];
        dfs(V);
        sb.append('\n');

        // 5. bfs
        visited = new boolean[N+1];
        bfs(V);

        // 6. 출력하기
        System.out.println(sb);
    }

    public static void dfs(int vertex) {
        visited[vertex] = true;
        sb.append(vertex).append(' ');

        for (int i = 0; i < adj[vertex].size(); i++) {
            int nextVertex = adj[vertex].get(i);

            if (!visited[nextVertex]) {
                dfs(nextVertex);
            }
        }
    }

    public static void bfs(int vertex) {
        queue = new LinkedList<>();
        queue.add(vertex);
        visited[vertex] = true;

        while(!queue.isEmpty()) {
            vertex = queue.poll();
            sb.append(vertex).append(' ');

            for (int i = 0; i < adj[vertex].size(); i++) {
                int nextVertex = adj[vertex].get(i);

                if (!visited[nextVertex]) {
                    queue.add(nextVertex);
                    visited[nextVertex] = true;
                }
            }
        }
    }
}
