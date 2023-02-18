package Baekjoon;

import java.io.*;

public class p2609 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        // 1. 입력받기
        String[] inputs = br.readLine().split(" ");
        int x = Integer.parseInt(inputs[0]);
        int y = Integer.parseInt(inputs[1]);

        // 2. 최대공약수 구하기 (유클리드 호제법)
        int gcd = getGcd(Math.max(x, y), Math.min(x, y));
        sb.append(gcd).append('\n');

        // 3. 최소공배수 구하기
        int lcm = (x/gcd) * (y/gcd) * gcd;
        sb.append(lcm).append('\n');

        // 4. 출력하기
        System.out.print(sb);
    }

    public static int getGcd(int a, int b) {
        return b>0 ? getGcd(b, a%b) : a;
    }
}
