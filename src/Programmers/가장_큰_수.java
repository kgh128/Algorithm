package Programmers;

import java.util.Arrays;

public class 가장_큰_수 {
    public String solution(int[] numbers) {
        String[] strings = new String[numbers.length];
        StringBuilder answer = new StringBuilder();

        // 1. int -> String으로 변환
        for (int i = 0; i < numbers.length; i++) {
            strings[i] = Integer.toString(numbers[i]);
        }

        // 2. 정렬 기준을 새로 만들어서 strings 정렬
        Arrays.sort(strings, (str1, str2)
                -> Integer.parseInt(str2 + str1) - Integer.parseInt(str1 + str2));

        // 3. 0이 맨 앞인 경우 (0만 있는 경우)
        if (strings[0].equals("0")) return "0";

        // 4. strings를 앞에서부터 이어붙이기
        for (String str: strings) {
            answer.append(str);
        }

        return answer.toString();
    }
}
