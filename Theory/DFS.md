# Depth-First Search

## DFS 순서
   ![DFS.png](/Theory/Image/DFS.png).
   
    1. root 노드 방문
    2. 그 노드의 모든 descendant 노드들을 차례로 방문
       - 보통 left -> right
       - preorder tree traversal



## DFS 기본 알고리즘

- 재귀함수나 스택으로 구현
  
    > 1. 탐색 시작 노드를 스택에 삽입하고 방문 처리
    > 2. 스택의 최상단 노드에 방문하지 않은 인접한 노드가 하나라도 있으면 그 노드를 스택에 넣고 방문처리
    > 3. 2번의 과정을 수행할 수 없을 때까지 반복

- 노드 방문 시 방문 여부를 반드시 검사
    > 아니면 무한 루프
- 깊이 제한(depth bound) 사용
    > 아니면 탐색 과정이 시작 노드에서 계속 깊어질 수 있음.
- 깊이 제한에 도달할 때까지 목표 노드가 발견되지 않으면 최근에 첨가된 노드의 부모 노드로 돌아와서(backtracking), 부모노드에 이전과는 다른 동작자를 적용하여 새로운 자식 노드를 생성

```cpp
// 코드 참고 : https://github.com/ndb796/python-for-coding-test 

#include <iostream>
#include <vector>
using namespace std;

// index 0은 사용하지 않음으로 배열을 하나 더 추가
bool visited[9]; 
vector<int> graph[9];

void dfs(int x)
{
	visited[x] = true;
	cout << x << " ";
	for (int i = 0; i < graph[x].size(); i++) // 인접한 노드 사이즈만큼 탐색
	{
		int y = graph[x][i];
		if (!visited[y]) // 방문하지 않았으면 즉 visited가 False일 때 not을 해주면 True가 되므로 아래 dfs 실행
            		dfs(y); // 재귀적으로 방문
	}
}

int main(void)
{
    /* 위 그래프와 동일하게 정의 */
    graph[1].push_back(2);
    graph[1].push_back(3);
    graph[1].push_back(8);

    graph[2].push_back(1);
    graph[2].push_back(7);

    graph[3].push_back(1);
    graph[3].push_back(4);
    graph[3].push_back(5);

    graph[4].push_back(3);
    graph[4].push_back(5);

    graph[5].push_back(3);
    graph[5].push_back(4);

    graph[6].push_back(7);

    graph[7].push_back(2);
    graph[7].push_back(6);
    graph[7].push_back(8);

    graph[8].push_back(1);
    graph[8].push_back(7);

    dfs(1);
}
```



## DFS 장점과 단점

- 장점
  
    1. 단지 현 경로상의 노드만을 기억하면 되므로 저장공간의 수요가 적다.
    2. 목표 노드가 깊은 단계에 있을 경우 해를 빨리 구할 수 있다.

- 단점

    1. 해가 없는 경로에 깊이 빠질 가능성이 있다.
        > 그러므로 미리 지정한 임의의 깊이까지만 탐색하고, 목표 노드를 발견하지 못하면 다음 경로를 따라 탐색하는 방법이 유용할 수 있다.
    2. 얻어진 해가 최단 경로가 된다는 보장이 없다.
        > 목표에 이르는 경로가 다수인 문제에 대해 DFS는 해에 다다르면 탐색을 끝내버리므로, 이때 얻어진 해는 최적이 아닐 수도 있다.



[출처](https://better-tomorrow.tistory.com/entry/DFS-BFS-%EC%9D%B4%ED%95%B4%ED%95%98%EA%B8%B0).
