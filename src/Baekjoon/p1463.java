package Baekjoon;

import java.io.*;

/*
[문제 풀이1] - dp: 상대적으로 시간이 걸리고, 메모리를 더 씀.
1. dp[X/2] + 1(X/2 연산 한 번 수행), dp[X/3] + 1(X/3 연산 한 번 수행), dp[X-1] + 1(X-1 연산 한 번 수행) 중 최솟값 구하기
2. 그 최솟값이 X를 1로 만드는데 사용하는 연산 횟수의 최솟값(dp[X])
*/

public class p1463 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        int[] dp = new int[N+1];

        // 2. 동적 프로그래밍 수행
        for (int x = 2; x < N+1; x++) {
            dp[x] = dp[x-1] + 1;
            if (x%2 == 0) dp[x] = Math.min(dp[x], dp[x/2] + 1);
            if (x%3 == 0) dp[x] = Math.min(dp[x], dp[x/3] + 1);
        }

        // 3. 출력하기
        System.out.println(dp[N]);
    }
}

/*
[문제 풀이2] - 재귀함수: 상대적으로 시간이 덜 걸리고, 메모리도 덜 씀.
1. 재귀함수를 돌면서 x/2를 했을 때의 count 값과 x/3을 했을 때의 count 값을 구하고, 그 중 최솟값을 반환한다.

1-1) x/2에 대해서 재귀함수를 호출했을 때 전달하는 count 값
     count(지금까지 누적 횟수) + 1(x/2 연산 수행) + x%2(x-1 연산 수행 횟수)
     -> 2로 나누어떨어지지 않는 x의 경우 2의 배수가 되기 위해서 x-1 연산을 수행해야 한다.
     -> 그 수행 횟수는 x를 2로 나눈 나머지이다.
     ex) x = 7이면 2의 배수가 되기 위해서 x-1 연산을 한 번(7%2 = 1) 수행하여 6으로 만들어야 한다.

1-2) x/3에 대해서 재귀함수를 호출했을 때 전달하는 count 값
     count(지금까지 누적 횟수) + 1(x/3 연산 수행) + x%3(x-1 연산 수행 횟수)
     -> 3으로 나누어떨어지지 않는 x의 경우 3의 배수가 되기 위해서 x-1 연산을 수행해야 한다.
     -> 그 수행 횟수는 x를 3로 나눈 나머지이다.
     ex) x = 10이면 3의 배수가 되기 위해서 x-1 연산을 한 번(10%3 = 1) 수행하여 9로 만들어야 한다.

2. x<2인 경우 지금까지 누적된 count 값을 반환 - 재귀 종료 조건
+) x=2인 경우 x/3의 값이 0이 나옴. -> x==1을 종료 조건으로 하면 계속 재귀 호출됨. (스택 오버플로우)
   x = x/2 = 2/2 = 1: count + 1 + 2%2 = count + 1
   x = x/3 = 2/3 = 0: count + 1 + 2%3 = count + 1 + 2 = count + 3
   위의 식처럼 당연히 x/2를 수행했을 때의 count 값이 더 작으므로 x 값이 0인 경우는 걸러져서 이대로 코드 짜도 상관X
*/

/*
public class p1463 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());

        // 2. 재귀함수 실행하여 연산 횟수의 최솟값 구하기
        System.out.println(minCount(N, 0));
    }

    public static int minCount(int x, int count) {
        if (x < 2) return count;

        return Math.min(minCount(x/2, count + 1 + x%2), minCount(x/3, count + 1 + x%3));
    }
}
*/