package Baekjoon;

import java.io.*;

public class p13458 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        // 1. N 입력받기
        String line1 = br.readLine();
        int N = Integer.parseInt(line1);

        // 2. A 배열 입력받기
        String[] line2 = br.readLine().split(" ");
        int[] A = new int[N];

        for (int i = 0; i < N; i++) {
            A[i] = Integer.parseInt(line2[i]);
        }

        // 3. B, C 입력받기
        String[] line3 = br.readLine().split(" ");
        int B = Integer.parseInt(line3[0]);
        int C = Integer.parseInt(line3[1]);

        // 4. 필요한 감독관 수의 최솟값 구하기
        long minSupervisors = N;

        for (int i = 0; i < N; i++) {
            // 총감독관으로만 감독이 가능하면 부감독관 필요X
            if (A[i] <= B) {
                continue;
            }

            // 부감독관 수 구하기
            A[i] -= B;  // 총감독관이 감시할 수 있는 응시자 수 빼기

            if (A[i] % C == 0) {
                minSupervisors += A[i] / C;
            }
            else {
                minSupervisors += A[i] / C + 1;
            }
        }

        bw.write(String.valueOf(minSupervisors));

        bw.flush();
        bw.close();
    }
}
