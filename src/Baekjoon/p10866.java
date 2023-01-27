package Baekjoon;

import java.io.*;

public class p10866 {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        StringBuilder bw = new StringBuilder();

        int N = Integer.parseInt(br.readLine());
        Deque deque = new Deque();

        for (int i = 0; i < N; i++) {
            String[] command = br.readLine().split(" ");

            if (command[0].equals("push_front")) {
                int x = Integer.parseInt(command[1]);
                deque.pushFront(x);
            }
            else if (command[0].equals("push_back")) {
                int x = Integer.parseInt(command[1]);
                deque.pushBack(x);
            }
            else if (command[0].equals("pop_front")) {
                bw.append(deque.popFront()).append('\n');
            }
            else if (command[0].equals("pop_back")) {
                bw.append(deque.popBack()).append('\n');
            }
            else if (command[0].equals("size")) {
                bw.append(deque.size()).append('\n');
            }
            else if (command[0].equals("empty")) {
                bw.append(deque.empty()).append('\n');
            }
            else if (command[0].equals("front")) {
                bw.append(deque.front()).append('\n');
            }
            else if (command[0].equals("back")) {
                bw.append(deque.back()).append('\n');
            }
        }

        System.out.print(bw);
    }

    private static class Deque {
        private int[] deque;
        private int frontIndex;
        private int backIndex;

        public Deque() {
            deque = new int[20000];
            frontIndex = 10000;
            backIndex = 9999;
        }

        public void pushFront(int x) {
            // 앞으로 한 칸 이동 (-1)
            deque[--frontIndex] = x;
        }

        public void pushBack(int x) {
            // 뒤로 한 칸 이동 (+1)
            deque[++backIndex] = x;
        }

        public int popFront() {
            if (empty() == 1) {
                return -1;
            }
            // 뒤로 한 칸 이동 (+1)
            return deque[frontIndex++];
        }

        public int popBack() {
            if (empty() == 1) {
                return -1;
            }
            // 앞으로 한 칸 이동 (-1)
            return deque[backIndex--];
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
            return deque[frontIndex];
        }

        public int back() {
            if (empty() == 1) {
                return -1;
            }
            return deque[backIndex];
        }
    }
}
