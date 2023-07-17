#include <iostream>
#include <string>

using namespace std;


// 전역변수 선언
int N, X;

// 스택 클래스 구현
class Stack {
    private:
        int container[10000];
        int topIdx = 0;
    
    public:
      void push(int X) {
        container[topIdx++] = X;
      } 

      int pop() {
        if (topIdx == 0) return -1;
        else return container[--topIdx];
      } 

      int size() {
        return topIdx;
      }

      int empty() {
        if (topIdx == 0) return 1;
        else return 0;
      }

      int top() {
        if (topIdx == 0) return -1;
        else return container[topIdx-1];
      }
};

int main(void) {
    // 0. 입출력 시간 줄이기
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    // 1. 변수 선언
    string cmd = "";
    Stack stack = Stack();

    // 2. 입력받아서 명령 처리
    cin >> N;

    for (int i = 0; i < N; i++) {
        cin >> cmd;

        if (cmd == "push") {
            cin >> X;
            stack.push(X);
        }
        else if (cmd == "pop") {
            cout << stack.pop() << '\n';
        }
        else if (cmd == "size") {
            cout << stack.size() << '\n';
        }
        else if (cmd == "empty") {
            cout << stack.empty() << '\n';
        }
        else if (cmd == "top") {
            cout << stack.top() << '\n';
        }
    }
}