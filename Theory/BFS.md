# Breadth-First Serach

## BFS 순서
   ![BFS.png](/Theory/Image/BFS.png).

    1. root 노드를 먼저 검색
    2. 단계1에 있는 모든 노드들을 검색
        - left -> right 
    3. level 2에 있는 모든 노드들을 검색
        - left -> right 
    .
    .
    .
    n+1. 단계n(가장 아래)에 있는 모든 노드들을 검색
        - leaves 모두 검색
        - left -> right   
  


## BFS 기본 알고리즘

- 큐를 사용해서 구현
  
    > 1. 탐색 시작 노드를 큐에 삽입하고 방문 처리
    > 2. 큐에서 노드를 꺼낸 뒤에 해당 노드의 인접 노드 중 방문하지 않은 노드를 모두 큐에 삽입 후 방문 처리
    > 3. 2번 반복

- 노드를 방문하고, 큐에 넣는다.
    > 해당 노드의 자식 노드들을 방문할 것이다.
- 노드를 큐에서 꺼낸다.
    > 해당 노드의 자식 노드들을 실제로 방문한다.
- 해당 노드의 자식 노드들을 한 번에 방문한다.
    > 단계 별로 방문한다.

```cpp
// 코드 참고 : https://github.com/ndb796/python-for-coding-test

#include <iostream>
#include <vector>
#include <queue>

using namespace std;

bool visited[9];
vector<int> graph[9];

// BFS 함수 정의
void bfs(int start) {
    queue<int> q;
    q.push(start); // 첫 노드를 queue에 삽입
    visited[start] = true; // 첫 노드를 방문 처리

    // 큐가 빌 때까지 반복
    while (!q.empty()) {
        // 큐에서 하나의 원소를 뽑아 출력
        int x = q.front();
        q.pop();
        cout << x << ' ';
        // 해당 원소와 연결된, 아직 방문하지 않은 원소들을 큐에 삽입
        for (int i = 0; i < graph[x].size(); i++) {
            int y = graph[x][i];
            if (!visited[y]) {
                q.push(y);
                visited[y] = true;
            }
        }
    }
}

int main(void) {
    // 노드 1에 연결된 노드 정보 저장 
    graph[1].push_back(2);
    graph[1].push_back(3);
    graph[1].push_back(8);

    // 노드 2에 연결된 노드 정보 저장 
    graph[2].push_back(1);
    graph[2].push_back(7);

    // 노드 3에 연결된 노드 정보 저장 
    graph[3].push_back(1);
    graph[3].push_back(4);
    graph[3].push_back(5);

    // 노드 4에 연결된 노드 정보 저장 
    graph[4].push_back(3);
    graph[4].push_back(5);

    // 노드 5에 연결된 노드 정보 저장 
    graph[5].push_back(3);
    graph[5].push_back(4);

    // 노드 6에 연결된 노드 정보 저장 
    graph[6].push_back(7);

    // 노드 7에 연결된 노드 정보 저장 
    graph[7].push_back(2);
    graph[7].push_back(6);
    graph[7].push_back(8);

    // 노드 8에 연결된 노드 정보 저장 
    graph[8].push_back(1);
    graph[8].push_back(7);

    bfs(1);
}
```


## BFS의 장점과 단점

- 장점
  
    1. 출발 노드에서 목표 노드까지의 최단 길이 경로를 보장
    > 특정 조건의 최단 경로 알고리즘을 계산할 때 사용

- 단점
  
    1. 경로가 매우 길 경우에는 탐색 branch가 급격히 증가함에 따라 많은 기억 공간을 필요로 함.
    2. 해가 존재하지 않는다면 유한 그래프의 경우, 모든 그래프를 탐색한 후에 실패로 끝남.
    3. 무한 그래프의 경우에는 결코 해를 찾지도 못하고, 끝내지도 못함.



[출처](https://better-tomorrow.tistory.com/entry/DFS-BFS-%EC%9D%B4%ED%95%B4%ED%95%98%EA%B8%B0).
