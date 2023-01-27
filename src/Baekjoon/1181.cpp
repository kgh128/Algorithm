#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int N;                  // 단어 개수
vector<string> w_list;  // 단어 리스트

bool compare(string a, string b) {
    // 2. 길이가 같으면 사전 순으로
    if (a.length() == b.length()) {
        return a.compare(b) < 0;
    }

    // 1. 길이가 짧은 것부터
    return a.length() < b.length();
} 

int main(void) {
    string buf = "";

    cin >> N;

    for (int i = 0; i < N; i++) {
        cin >> buf;
        w_list.push_back(buf);
    }

    sort(w_list.begin(), w_list.end(), compare); // 정렬
    w_list.erase(unique(w_list.begin(), w_list.end()), w_list.end()); // 중복 단어 제거

    for (int i = 0; i < w_list.size(); i++) {
        cout << w_list[i] << '\n';
    }
}