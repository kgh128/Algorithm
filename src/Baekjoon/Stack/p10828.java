package Baekjoon.Stack;

import java.io.*;

public class p10828 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        int N = Integer.parseInt(br.readLine());
        Stack stack = new Stack();

        for (int i = 0; i < N; i++) {
            String[] command = br.readLine().split(" ");

            if (command[0].equals("push")) {
                int X = Integer.parseInt(command[1]);
                stack.push(X);
            }
            else if (command[0].equals("pop")) {
                bw.append(stack.pop()).append('\n');
            }
            else if (command[0].equals("size")) {
                bw.append(stack.size()).append('\n');
            }
            else if (command[0].equals("empty")) {
                bw.append(stack.empty()).append('\n');
            }
            else if (command[0].equals("top")) {
                bw.append(stack.top()).append('\n');
            }
        }

        System.out.print(bw);
    }

    private static class Stack {
        private int[] stack;
        private int top;

        public Stack() {
            stack = new int[10000];
            top = -1;
        }

        public void push(int X) {
            stack[++top] = X;
        }

        public int pop() {
            if (empty() == 1) {
                return -1;
            }
            return stack[top--];
        }

        public int size() {
            return top+1;
        }

        public int empty() {
            if (top < 0) {
                return 1;
            }
            return 0;
        }

        public int top() {
            if (empty() == 1) {
                return -1;
            }
            return stack[top];
        }
    }
}
