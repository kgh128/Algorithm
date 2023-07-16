#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int N;
vector<int> v;

int main(void) {
    // 1. 입력받기
    cin >> N;

    for (int i = 0; i< N; i++) {
        int tmp;
        cin >> tmp;
        v.push_back(tmp);
    }

    // 2. 오름차순 정렬
    sort(v.begin(), v.end());

    // 3. 출력하기
    for (int i = 0; i < N; i++) {
        cout << v[i] << '\n';
    }
}