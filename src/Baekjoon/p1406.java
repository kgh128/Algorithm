package Baekjoon;

import java.io.*;
import java.util.*;

public class p1406 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        // 0. 커서를 기준으로 문자열을 왼쪽, 오른쪽으로 쪼개서 각각의 스택에 저장할 것
        //    (커서는 각 스택의 top)
        Stack<Character> leftStack = new Stack<>();     // 커서 왼쪽의 문자열
        Stack<Character> rightStack = new Stack<>();    // 커서 오른쪽의 문자열

        // 1. 편집기에 입력된 문자열을 leftStack에 넣기
        //    (처음에는 커서가 문자열의 끝에 있으므로)
        String str = br.readLine();

        for (int i = 0; i < str.length(); i++) {
            leftStack.push(str.charAt(i));
        }

        // 2. 명령어 수행
        int M = Integer.parseInt(br.readLine());

        for (int i = 0; i < M; i++) {
            String[] command = br.readLine().split(" ");

            switch(command[0]) {
                case "L": // leftStack의 top 문자 -> rightStack
                    if (!leftStack.isEmpty()) {
                        rightStack.push(leftStack.pop());
                    }
                    break;

                case "D": // rightStack의 top 문자 -> leftStack
                    if (!rightStack.isEmpty()) {
                        leftStack.push(rightStack.pop());
                    }
                    break;

                case "B": // leftStack의 top 문자 삭제
                    if (!leftStack.isEmpty()) {
                        leftStack.pop();
                    }
                    break;

                case "P": // leftStack에 문자 추가
                    leftStack.push(command[1].charAt(0));
                    break;
            }
        }

        // 3. leftStack과 rightStack 합쳐서 최종 문자열 만들기
        while(!leftStack.isEmpty()) {
            rightStack.push(leftStack.pop());
        }

        while(!rightStack.isEmpty()) {
            bw.append(rightStack.pop());
        }

        System.out.println(bw);
    }
}
