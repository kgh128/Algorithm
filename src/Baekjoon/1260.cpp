#include <iostream>
#include <vector>
#include <queue>
#include <algorithm>
using namespace std;

bool dfs_visited[1001];
bool bfs_visited[1001];

vector<vector <int>> graph;

void dfs(int x)
{
    // 노드 x 방문
    dfs_visited[x] = true;
    cout << x << " ";

    // x의 인접한 노드 개수만큼 탐색
    for (int i = 0; i < graph[x].size(); i++) {
        int y = graph[x][i];

        // 노드 y를 방문하지 않았으면 방문
        if (!dfs_visited[y]) {
            dfs(y);
        }
    }
}

void bfs(int start) {
    queue<int> q;
    q.push(start);
    bfs_visited[start] = true;

    // 큐가 빌 때까지 반복
    while(!q.empty()) {
        int x = q.front();
        q.pop();
        cout << x << " ";

        // x의 인접한 노드들 중에서 아직 방문하지 않은 노드들을 큐에 삽입
        for (int i = 0; i < graph[x].size(); i++) {
            int y = graph[x][i];
            if (!bfs_visited[y]) {
                q.push(y);
                bfs_visited[y] = true;
            }
        }
    }
}

int main(void) {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);

    int N, M, start;
    int x, y;

    // 그래프 만들기
    // 1. 노드 리스트
    cin >> N >> M >> start;
    graph.resize(N+1);

    // 2. 각 노드의 인접 노드 리스트
    for (int i = 0; i < M; i++) {
        cin >> x >> y;
        graph[x].push_back(y);
        graph[y].push_back(x);
    }

    // 3. 각 노드의 인접 노드 리스트 정렬
    for (int i = 1; i < N+1; i++) {
        sort(graph[i].begin(), graph[i].end());
    }

    // DFS 수행 결과
    dfs(start);
    cout << "\n";

    // BFS 수행 결과
    bfs(start);
    cout << "\n";
}