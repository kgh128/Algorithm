package Baekjoon;

import java.io.*;
import java.util.*;

public class p11650 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        ArrayList<Pair> pairs = new ArrayList<>(N);

        for (int i = 0; i < N; i++) {
            String[] input = br.readLine().split(" ");

            int x = Integer.parseInt(input[0]);
            int y = Integer.parseInt(input[1]);

            pairs.add(new Pair(x, y));
        }

        // 2. 정렬하기
        pairs.sort((a, b) -> {
            if (a.getX() == b.getX()) return a.getY() - b.getY();
            else return a.getX() - b.getX();
        });

        // 3. 출력하기
        for (int i = 0; i < N; i++) {
            Pair pair = pairs.get(i);
            bw.append(pair.getX()).append(' ').append(pair.getY()).append('\n');
        }
        System.out.print(bw);
    }

    private static class Pair {
        private final int x;
        private final int y;

        public Pair(int x, int y) {
            this.x = x;
            this.y = y;
        }

        public int getX() {
            return x;
        }

        public int getY() {
            return y;
        }
    }
}
