#include <iostream>
#include <string>

using namespace std;


// 전역변수 선언
int N, X;

// 큐 클래스 구현
class Queue {
    private:
        int container[10000];
        int frontIdx = 0; 
        int backIdx = 0;
    
    public:
        void push(int X) {
            container[backIdx++] = X;
        }

        int pop() {
            if (frontIdx == backIdx) return -1;
            else return container[frontIdx++];
        }

        int size() {
            return backIdx - frontIdx;
        }

        int empty() {
            if (frontIdx == backIdx) return 1;
            else return 0;
        }

        int front() {
            if (frontIdx == backIdx) return -1;
            else return container[frontIdx];
        }

        int back() {
            if (frontIdx == backIdx) return -1;
            else return container[backIdx-1];
        }
};

int main(void) {
    // 0. 입출력 시간 줄이기
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    // 1. 변수 선언
    string cmd = "";
    Queue queue = Queue();

    // 2. 입력받아서 명령 처리
    cin >> N;

    for (int i = 0; i < N; i++) {
        cin >> cmd;

        if (cmd == "push") {
            cin >> X;
            queue.push(X);
        }
        else if (cmd == "pop") {
            cout << queue.pop() << '\n';
        }
        else if (cmd == "size") {
            cout << queue.size() << '\n';
        }
        else if (cmd == "empty") {
            cout << queue.empty() << '\n';
        }
        else if (cmd == "front") {
            cout << queue.front() << '\n';
        }
        else if (cmd == "back") {
            cout << queue.back() << '\n';
        }
    }
}