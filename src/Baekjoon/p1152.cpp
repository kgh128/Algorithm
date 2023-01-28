#include <iostream>
using namespace std;

int main(void) {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);

    int n = 0;              //단어의 개수
    char buf[1000001];      // 문자열을 받을 버퍼

    // 1. 문자열 입력 받기
    // 문자열의 종결 문자는 널문자(\0)
    cin.getline(buf, 1000001); 

    // 1. 단어의 개수 세기
    for (int i = 0; i < 1000000; i++) {
        // 종결 문자 나오면 종료
        if (buf[i] == '\0') break;

        // i = 0일 때 처리 (문자열의 처음 문자)
        // (i-1 접근이 불가능하므로 따로 처리)
        // 알파벳이면 단어의 개수 + 1 (단어의 시작 부분이므로)
        // 공백이면 패스
        if (i == 0) {
            if (buf[i] != ' ') n++;
            continue;
        }

        // 이전 문자가 공백이고, 현재 문자가 알파벳이면 단어의 시작 부분
        // 따라서 단어의 개수 + 1
        if (buf[i-1] == ' ' && buf[i] != ' ') {
            n++;
        }
    }

    // 2. 단어의 개수 출력
    cout << n << '\n';

}