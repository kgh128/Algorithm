#include <iostream>
#include <limits>
using namespace std;

int W, H;                   // 지도의 너비와 높이
int islands;                // 섬의 개수
int map[52][52];            // 지도
bool visited[52][52];       // 방문 여부

void dfs(int i, int j) {
    visited[i][j] = true;

    if (map[i-1][j-1] == 1 && visited[i-1][j-1] == false) {
        dfs(i-1, j-1);
    }

    if (map[i-1][j] == 1 && visited[i-1][j] == false) {
        dfs(i-1, j);
    }

    if (map[i-1][j+1] == 1 && visited[i-1][j+1] == false) {
        dfs(i-1, j+1);
    }

    if (map[i][j-1] == 1 && visited[i][j-1] == false) {
        dfs(i, j-1);
    }

    if (map[i][j+1] == 1 && visited[i][j+1] == false) {
        dfs(i, j+1);
    }

    if (map[i+1][j-1] == 1 && visited[i+1][j-1] == false) {
        dfs(i+1, j-1);
    }

    if (map[i+1][j] == 1 && visited[i+1][j] == false) {
        dfs(i+1, j);
    }

    if (map[i+1][j+1] == 1 && visited[i+1][j+1] == false) {
        dfs(i+1, j+1);
    }
    
    return;
}

int main(void) {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    
    do {
        cin >> W >> H;

        // 1. 지도를 0으로 감싸기
        for (int i = 0; i <= W+1; i++) {
            map[0][i] = 0; // 맨 윗줄
            map[H+1][i] = 0; // 맨 아랫줄
        }
        
        for (int i = 1; i <= H; i++) {
            map[i][0] = 0; // 맨 왼쪽 줄
            map[i][W+1] = 0; // 맨 오른쪽 줄
        }

        // 2. 지도 채워넣기
        for (int i = 1; i <= H; i++) {
            for (int j = 1; j <= W; j++) {
                cin >> map[i][j];
            }
        }

        // 3. DFS로 섬 찾기
        for (int i = 1; i <= H; i++) {
            for (int j = 1; j <= W; j++) {
                if (map[i][j] == 1 && visited[i][j] == false) {
                    dfs(i, j);
                    islands++;
                }
            }
        }

        // 4. 정답 출력
        if (!(W == 0 && H == 0)) {
            cout << islands << "\n";
        }

        // 5. 초기화
        islands = 0;

        for (int i = 1; i <= H; i++) {
            for (int j = 1; j <= W; j++) {
                visited[i][j] = false;
            }
        }
    } while(!(W == 0 && H == 0));
}