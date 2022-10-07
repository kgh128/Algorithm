# Depth-First Search

## DFS 순서
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

```cpp
void dfs(node v) {
    node u;

    visit v;
    for (each child u of v) {
        dfs(u);
    }
}
```