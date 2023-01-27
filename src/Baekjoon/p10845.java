package Baekjoon;

import java.io.*;

public class p10845 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        int N = Integer.parseInt(br.readLine());
        Queue queue = new Queue();

        for (int i = 0; i < N; i++) {
            String[] command = br.readLine().split(" ");

            if (command[0].equals("push")) {
                int x = Integer.parseInt(command[1]);
                queue.push(x);
            }
            else if (command[0].equals("pop")) {
                bw.append(queue.pop()).append('\n');
            }
            else if (command[0].equals("size")) {
                bw.append(queue.size()).append('\n');
            }
            else if (command[0].equals("empty")) {
                bw.append(queue.empty()).append('\n');
            }
            else if (command[0].equals("front")) {
                bw.append(queue.front()).append('\n');
            }
            else if (command[0].equals("back")) {
                bw.append(queue.back()).append('\n');
            }
        }

        System.out.print(bw);
    }

    private static class Queue {
        private int[] queue;
        private int frontIndex;
        private int backIndex;

        public Queue() {
            queue = new int[10000];
            frontIndex = 0;
            backIndex = -1;
        }

        public void push(int x) {
            queue[++backIndex] = x;
        }

        public int pop() {
            if (empty() == 1) {
                return -1;
            }
            return queue[frontIndex++];
        }

        public int size() {
            return backIndex - frontIndex + 1;
        }

        public int empty() {
            if (backIndex < frontIndex) {
                return 1;
            }
            return 0;
        }

        public int front() {
            if (empty() == 1) {
                return -1;
            }
            return queue[frontIndex];
        }

        public int back() {
            if (empty() == 1) {
                return -1;
            }
            return queue[backIndex];
        }
    }
}
