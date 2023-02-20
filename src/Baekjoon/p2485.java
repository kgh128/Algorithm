package Baekjoon;

import java.io.*;

public class p2485 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[] tree = new int[N];

        for (int i = 0; i < N; i++) {
            tree[i] = Integer.parseInt(br.readLine());
        }

        // 2. 가로수 간격의 최대 공약수 구하기
        int gcd = tree[1] - tree[0];

        for (int i = 2; i < N; i++) {
            gcd = getGcd(tree[i] - tree[i-1], gcd);
        }

        // 3. 새로 심어야 하는 가로수의 개수 구하기
        int count = 0;

        for (int i = 1; i < N; i++) {
            count += (tree[i] - tree[i-1]) / gcd - 1;
        }

        System.out.println(count);

    }

    public static int getGcd(int a, int b) {
        return b != 0 ? getGcd(b, a%b) : a;
    }
}
