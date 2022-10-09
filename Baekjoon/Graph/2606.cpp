#include <iostream>
#include <vector>
using namespace std;

int V, E;                   // 컴퓨터(정점)의 수, 연결쌍(간선)의 수
int infection;              // 감염된 컴퓨터의 수
bool visited[101];          // 컴퓨터 방문 여부
vector<vector <int>> graph; // 컴퓨터 네트워크 그래프

void dfs(int x) {
    visited[x] = true;

    for (int i = 0; i < graph[x].size(); i++) {
        int y = graph[x][i];
        
        if (!visited[y]) {
            infection++;    // 감염된 컴퓨터의 수 + 1
            dfs(y);
        }
    }
}

int main(void) {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);

    int x, y;   // 임시 변수

    // 1. 컴퓨터 네트워크 그래프 만들기
    cin >> V >> E;
    graph.resize(V+1);  // 컴퓨터들의 리스트 만들기 (0번은 사용X)

    // 각 컴퓨터에 연결된 컴퓨터 리스트 만들기 (인접 행렬 만들기)
    for (int i = 0; i < E; i++) {
        cin >> x >> y;
        graph[x].push_back(y);
        graph[y].push_back(x);
    }

    // 2. 1번 컴퓨터에 대해 DFS
    // (1번 컴퓨터와 연결된 컴퓨터들 찾기)
    dfs(1);

    // 3. 감염된 컴퓨터의 수 출력
    cout << infection << "\n";
}