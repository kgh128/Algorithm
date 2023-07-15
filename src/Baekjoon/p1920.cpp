#include <iostream>
#include <algorithm>

using namespace std;

int N, M;
int A[100000];
int B[100000];

int main(){
    // 1. 입력받기
    cin >> N;

    for (int i = 0; i < N; i++) {
        cin >> A[i];
    }

    cin >> M;

    for (int i = 0; i < M; i++) {
        cin >> B[i];
    }

    // 2. 배열 A 오름차순 정렬
    sort(A, A+N);
    
    // 3. B[i]가 A에 안에 있는지 이분탐색
    for (int i = 0; i < M; i++) {
        if (binary_search(A, A+N, B[i])) {
            cout << 1 << '\n';
        }
        else {
            cout << 0 << '\n';
        }
    }
}