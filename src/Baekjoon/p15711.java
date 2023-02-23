package Baekjoon;

import java.io.*;
import java.util.*;

public class p15711 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder sb = new StringBuilder();

        // 1. 소수 판정해놓기 - 에라토스테네스의 체
        boolean[] isNotPrime = new boolean[2000001];
        ArrayList<Integer> primes = new ArrayList<>();

        isNotPrime[1] = true;

        for (int i = 2; i < 2000001; i++) {
            if (!isNotPrime[i]) {
                primes.add(i);

                for (int j = 2; i * j < 2000001; j++) {
                    isNotPrime[i * j] = true;
                }
            }
        }

        // 2. 테스트 케이스 실행
        int T = Integer.parseInt(br.readLine());

        for (int i = 0; i < T; i++) {
            // 3. 입력받기
            StringTokenizer st = new StringTokenizer(br.readLine());
            long A = Long.parseLong(st.nextToken());
            long B = Long.parseLong(st.nextToken());
            boolean possible = true;

            // 5. A+B가 홀수일 경우: 홀수+짝수로 쪼개짐.
            // 짝수인 소수는 2밖에 없으니 A+B-2가 소수인지 판별
            if ((A + B) % 2 == 1) {
                long rest = A + B - 2;

                // A+B-2가 에라토스테네스의 체 수행 범위에 있는 경우
                if (rest < 2000001) {
                    if (isNotPrime[(int)rest]) {
                        possible = false;
                    }
                }
                // A+B-2가 에라토스테네스의 체 수행 범위를 넘어가는 경우
                // 모든 합성수는 소수의 배수이므로 지금까지 구한 소수로 나누어 떨어지는지 확인
                else {
                    for (int prime: primes) {
                        if (rest % prime == 0) {
                            possible = false;
                            break;
                        }
                    }
                }
            }

            // 4. A+B가 짝수일 경우: 2보다 큰 짝수는 소수의 합으로 쪼개짐. (골드바흐의 추측)
            // 10^18 이하에서 성립하는데, A+B의 최댓값은 4 * 10^12이므로 사용할 수 있다.
            else if (A+B == 2) possible = false;

            // 6. 출력하기
            if (possible) sb.append("YES\n");
            else sb.append("NO\n");
        }
        System.out.print(sb);
    }
}

/*
[에라토스테네스의 체를 2000000까지만 수행하는 이유]
1. 자바에서 배열의 최대 크기는 int의 최댓값과 같은데, A+B의 최댓값인 4 * 10^12는 int 범위를 넘어간다.
   따라서 4 * 10^12 + 1의 크기를 가지는 boolean 배열을 만들 수 없다.
   boolean[] isNotPrime = new boolean[4000000000001]; // 불가능

2. 4 * 10^12 이하의 합성수는 반드시 2000000 이하의 소수를 약수로 가지고 있다.
   (합성수) = (약수1) * (약수2)          -> ex) 34 = 2 * 17
                                      -> ex) 16 = 2 * 8
   (약수1) <= (합성수의 제곱근) <= (약수2) -> ex) 2 <= (34의 제곱근인 5.xxx) <= 17
                                       -> ex) 2 <= (16의 제곱근인 4) <= 8

   위와 같은 규칙이 성립하므로 (합성수의 제곱근) 이상의 소수의 배수이면 당연히 (합성수의 제곱근) 이하의 소수의 배수일 수밖에 없다.
   따라서 (합성수의 제곱근) 이하의 소수의 배수인지만 확인하면 된다.
   -> (약수2)가 소수더라도 (약수2)까지 갈 필요 없이 (약수1)로 나누어 떨어질 것이고,
      (약수1)은 소수이거나, 자기보다 작은 소수의 배수일 것이다.
      이 소수는 (약수1)보다 작거나 같으므로 당연히 (합성수의 제곱근) 이하일 것이다.
      따라서 (합성수)는 (합성수의 제곱근) 이하의 소수의 배수이다.

   이 문제에서 A+B의 최댓값은 4 * 10^12인데, 이 수의 제곱근은 2000000이다.
   따라서 4 * 10^12는 2000000 이하의 소수의 배수일 것이므로, 4 * 10^12보다 작은 합성수들도 당연히 2000000 이하의 소수의 배수일 것이다.
   그러므로 에라토스테네스의 체로는 2000000 이하의 소수들을 구하고, A+B-2가 2000000보다 클 경우 2000000 이하의 소수로 나눠떨어지는지 확인한다.
   나눠떨어지면 합성수이고, 모든 2000000 이하의 소수들로 안나눠떨어지면 소수이다.
 */