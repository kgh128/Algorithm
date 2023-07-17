#include <iostream>

using namespace std;

// 주어지는 수는 1-10000의 자연수이므로
// 해당 자연수들의 개수를 세서 배열에 저장하는 형식으로 정렬. (계수 정렬 사용)
// 그러면 int 배열의 원소 개수가 최대 10001개가 필요하므로 약 40KB만 있으면 됨.
// (1-10000의 자연수를 바로 인덱스에 매칭하기 위해서는 10001개의 원소가 필요)
// -> 나중에 count 배열을 앞에서부터 출력하면 정렬한 결과가 출력됨.
// 10,000,000개의 데이터를 다 배열에 저장해서 정렬하는 형식은
// short형 배열로 만든다고 해도 20MB가 필요하므로, 8MB 메모리 제한에 걸림.


int N;
int count[10001]; // index: 입력으로 받은 수, value: 해당 수의 개수

int main(void) {
    // 0. 입출력 시간 줄이기
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    
    // 1. 입력받기
    cin >> N;

    for (int i = 0; i < N; i++) {
        int num;
        cin >> num;
        
        // 2. 입력받은 해당 자연수의 개수 세기
        count[num]++;
    }

    // 3. 각 자연수의 count 개수가 0이 될 때까지 출력하기
    // 1부터 10000까지 반복문을 시행하므로 출력되는 값은 정렬이 되어있음.
    for (int i = 1; i <= 10000; i++) {
        // 같은 수가 중복되어서 들어올 수 있기 때문에
        // 중복되어서 들어온만큼 다 출력해줘야 함.
        for (int j = count[i]; j > 0; j--) {
            cout << i << '\n';
        }
    }
}