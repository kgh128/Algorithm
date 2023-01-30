package Baekjoon;

import java.io.*;
import java.util.*;

public class p1181 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 1. 중복 제거해서 입력받기
        int N = Integer.parseInt(br.readLine());
        HashSet<String> wordSet = new HashSet<>(N);

        for (int i = 0; i < N; i++) {
            wordSet.add(br.readLine());
        }

        // 2. 길이가 짧은 순으로 정렬
        ArrayList<String> wordList = new ArrayList<>(wordSet);

        Comparator<String> comparator = (a, b) -> {
            if (a.length() == b.length()) {
                return a.compareTo(b);
            }
            return a.length() - b.length();
        };

        wordList.sort(comparator);

        // 3. 출력
        for (String word: wordList) {
            bw.append(word).append('\n');
        }
        System.out.print(bw);
    }
}
