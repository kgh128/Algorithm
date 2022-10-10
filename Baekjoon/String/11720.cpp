#include <iostream>
using namespace std;

int N;

int main(void) {
    int n;
    int sum = 0;

    // input
    cin >> N;
    for(int i=0; i<N; i++) {  
        scanf("%1d", &n);
        sum += n;
    }

    cout << sum;
}