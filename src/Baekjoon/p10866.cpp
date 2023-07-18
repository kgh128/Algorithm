#include <iostream>
#include <string>

using namespace std;


// 전역변수 선언
int N, X;

// Deque 클래스 구현
class Deque {
    private:
        int container[20000];
        int frontIdx = 9999;
        int backIdx = 10000;
    
    public:
        void push_front(int X) {
            container[frontIdx--] = X;
        }

        void push_back(int X) {
            container[backIdx++] = X;
        }

        int pop_front() {
            if (backIdx - frontIdx == 1) return -1;
            else return container[++frontIdx];
        }

        int pop_back() {
            if (backIdx - frontIdx == 1) return -1;
            else return container[--backIdx];
        }

        int size() {
            return backIdx - frontIdx - 1;
        }

        int empty() {
            if (backIdx - frontIdx == 1) return 1;
            else return 0;
        }

        int front() {
            if (backIdx - frontIdx == 1) return -1;
            else return container[frontIdx + 1];
        }

        int back() {
            if (backIdx - frontIdx == 1) return -1;
            else return container[backIdx - 1];
        }
};

int main(void) {
    // 0. 입출력 시간 줄이기
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    // 1. 변수 선언
    string cmd = "";
    Deque deque = Deque();

    // 2. 입력받아서 명령 처리
    cin >> N;

    for (int i = 0; i < N; i++) {
        cin >> cmd;

        if (cmd == "push_front") {
            cin >> X;
            deque.push_front(X);
        }
        else if (cmd == "push_back") {
            cin >> X;
            deque.push_back(X);
        }
        else if (cmd == "pop_front") {
            cout << deque.pop_front() << '\n';
        } 
        else if (cmd == "pop_back") {
            cout << deque.pop_back() << '\n';
        }
        else if (cmd == "size") {
            cout << deque.size() << '\n';
        }
        else if (cmd == "empty") {
            cout << deque.empty() << '\n';
        }
        else if (cmd == "front") {
            cout << deque.front() << '\n';
        }
        else if (cmd == "back") {
            cout << deque.back() << '\n';
        }
    }
}