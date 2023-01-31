package Baekjoon;

import java.io.*;
import java.util.HashMap;

public class p10816 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 입력받기
        int N = Integer.parseInt(br.readLine());
        String[] inputs = br.readLine().split(" ");
        HashMap<Integer, Integer> ownCards = new HashMap<>(N);

        for (int i = 0; i < N; i++) {
            int card = Integer.parseInt(inputs[i]);
            ownCards.put(card, ownCards.getOrDefault(card, 0) + 1);
        }

        int M = Integer.parseInt(br.readLine());
        String[] givenCards = br.readLine().split(" ");

        // 2. 검색하기
        for (String givenCard: givenCards) {
            int card = Integer.parseInt(givenCard);
            bw.append(ownCards.getOrDefault(card, 0)).append(' ');
        }

        // 3. 출력하기
        System.out.println(bw);
    }
}
