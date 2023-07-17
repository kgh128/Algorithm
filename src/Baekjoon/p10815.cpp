#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;


int N,M;
vector<int> card;
vector<int> target;

int main(void) {
    // 0. 입출력 시간 줄이기
    ios::sync_with_stdio(false);
    cin.tie(NULL);

    // 1. 입력받기
    int buf = 0;

    cin >> N;

    for (int i = 0; i < N; i++) {
        cin >> buf;
        card.push_back(buf);
    }

    cin >> M;

    for (int i = 0; i < M; i++) {
        cin >> buf;
        target.push_back(buf);
    }

    // 2. card 오름차순 정렬하기
    sort(card.begin(), card.end());

    // 3. target을 card 안에서 이분탐색
    for (int i = 0; i < M; i++) {
        if (binary_search(card.begin(), card.end(), target[i])) {
            cout << "1 ";
        }
        else {
            cout << "0 ";
        }
    }
}