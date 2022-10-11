#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int T;                      // 테스트 케이스 개수
int M, N, K;                // 밭의 가로길이, 세로길이, 배추의 개수
int farm[50][50];           // 밭의 지도
bool visited[50][50];       // 배추 방문 여부 
int worms;                  // 필요한 최소의 지렁이 개수
pair<int, int> buf[2500];     // 각 테스트 케이스의 배추 위치 저장할 임시 버퍼
                            // X: buf[i].second, Y: buf[i].first

// 상하좌우 방향
int dy[4] = {-1, 1, 0, 0};
int dx[4] = {0, 0, -1, 1};


void bfs(int start_y, int start_x) {
    queue<pair <int, int>> q;
    pair<int, int> v; // 기준 노드
    pair<int, int> n; // 연결 노드

    // 1. 시작 노드 방문
    v = make_pair(start_y, start_x);
    q.push(v);
    visited[v.first][v.second] = true;

    // 2. 큐가 빌 때까지 반복하면서 탐색
    while(!q.empty()) {
        // 큐에서 노드를 꺼내서 해당 노드와 연결된 노드들을 탐색
        v = q.front();
        q.pop();

        // 상하좌우의 노드 탐색
        for (int i = 0; i < 4; i++) {
            n.first = v.first + dy[i];
            n.second = v.second + dx[i];

            // 범위를 넘어가는 노드들은 패스
            if (n.first < 0 || n.first > N-1 || n.second < 0 || n.second > M-1) {
                continue;
            }

            // v에 인접한 노드들 중에서 아직 방문하지 않은 노드들을 큐에 삽입
            if (farm[n.first][n.second] == 1 && visited[n.first][n.second] == false){
                q.push(n);
                visited[n.first][n.second] = true;
            }
        }
    }
}

int main(void) {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    
    // 1. 테스트 케이스 개수 받기
    cin >> T;

    // 2. 각 테스트 케이스 수행
    for (int i = 0; i < T; i++) {
        cin >> M >> N >> K;

        // 3. 지도에 배추 위치 표시
        for (int j = 0; j < K; j++) {
            cin >> buf[j].second >> buf[j].first;
            farm[buf[j].first][buf[j].second] = 1;
        }

        // 4. bfs로 연결된 배추들 탐색
        // 탐색할 때 buf 배열을 이용하여 배추(1)들만 탐색 -> 이중 for문 사용안해도 됨.
        // (0인 곳은 아예 접근 X)
        // 상하좌우로 연결된 배추들 덩어리를 찾으면 지렁이 + 1
        for (int j = 0; j < K; j++) {
            int y = buf[j].first;
            int x = buf[j].second;

            if (visited[y][x] == false) {
                bfs(y, x);
                worms++;
            }
        }

        // 5. 지렁이 마리 수 출력
        cout << worms << '\n';

        // 6. 초기화
        worms = 0;
        for (int j = 0; j < K; j++) {
            farm[buf[j].first][buf[j].second] = 0;
            visited[buf[j].first][buf[j].second] = false;
        }
    }
}