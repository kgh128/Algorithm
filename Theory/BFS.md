# Breadth-First Serach

## BFS 순서

    1. root 노드를 먼저 검색
    2. level 1에 있는 모든 노드들을 검색
        - left -> right 
    3. level 2에 있는 모든 노드들을 검색
        - left -> right 
    .
    .
    .
    n+1. level n(가장 아래)에 있는 모든 노드들을 검색
        - leaves 모두 검색
        - left -> right   
  


## BFS 기본 알고리즘

- Queue를 사용!
    > 다음에 호출해야하는 순서를 Queue로 유지
- 노드를 방문하고, Queue에 넣는다.
    > 해당 노드의 child 노드들을 방문할 것이다.
- 노드를 Queue에서 꺼낸다.
    > 해당 노드의 child 노드들을 실제로 방문한다.
- 해당 노드의 child 노드들을 한 번에 방문한다.
    > level 별로 방문한다.

```cpp
void bfs(tree T) {
    queue Q;
    node u, v;

    initialize(Q);
    v = root of T

    visit v;
    enqueue(Q, v);
    // v를 Queue에 넣으면 
    // v의 child 노드들을 이후에 while문에서 방문 
    
    while(!empty(Q)) {
        dequeue(Q, v);
        for (each child u of v) {
            visit u;
            enqueue(Q, u);
            // u의 child 노드들을
            // 이후에 while문에서 방문
        }
    }
}
```