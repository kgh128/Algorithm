package Baekjoon;

import java.io.*;

public class p2661 {
    static int N;
    static StringBuilder sb;

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        // 1. 입력받기
        N = Integer.parseInt(br.readLine());
        sb = new StringBuilder();

        // 2. 백트래킹 하면서 좋은 수열 찾기
        sb.append(1); // 초기값은 가장 작은 수인 1
        getGoodSequential(1);

        // 3. 출력하기
        System.out.println(sb);
    }

    public static boolean getGoodSequential(int length) {
        // 인접한 부분 수열이 동일한지 확인 (현재 수열이 좋은 수열인지 확인)
        for (int i = 1; i <= length/2; i++) {
            if (sb.substring(length-i).equals(sb.substring(length-2*i, length-i))) {
                return false;
            }
        }

        // 정해진 길이의 좋은 수열을 구한 경우
        if (length == N) return true;

        // 다음 수 넣기 - 1, 2, 3
        for (int i = 1; i <= 3; i++) {
            sb.append(i);
            length++;

            if (getGoodSequential(length)) {
                return true;
            }

            sb.deleteCharAt(length-1);
            length--;
        }

        // 기본 리턴값 (아직 좋은 수열을 구하지 못함.)
        return false;
    }
}

/*
[현재 수열이 좋은 수열인지 확인하는 방법]
- 현재 수열에서 1 ~ length/2 길이의 부분 수열을 모두 확인
  (2번 이상 존재할 수 있는 부분 수열의 최대 길이는 length/2)

- 앞의 부분 수열들은 이미 검증이 끝난 것이므로 가장 끝에 있는 1 ~ length/2 길이의 부분 수열과
  해당 부분 수열 바로 앞에 있는 1 ~ length/2 길이의 부분 수열끼리만 비교
 */