package Baekjoon;

import java.io.*;
import java.util.ArrayList;

public class p2606 {
    static int com;
    static int pair;
    static int count;
    static ArrayList<Integer>[] adj;
    static boolean[] visited;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        com = Integer.parseInt(br.readLine());
        pair = Integer.parseInt(br.readLine());
        adj = new ArrayList[com+1];
        visited = new boolean[com+1];

        for (int i = 0; i < com+1; i++) {
            adj[i] = new ArrayList<>();
        }

        for (int i = 0; i < pair; i++) {
            String[] inputs = br.readLine().split(" ");
            int x = Integer.parseInt(inputs[0]);
            int y = Integer.parseInt(inputs[1]);

            adj[x].add(y);
            adj[y].add(x);
        }

        // 2. dfs
        dfs(1);
        System.out.println(count);
    }

    public static void dfs(int current) {
        visited[current] = true;

        for (int next: adj[current]) {
            if (!visited[next]) {
                count++;
                dfs(next);
            }
        }
    }
}
