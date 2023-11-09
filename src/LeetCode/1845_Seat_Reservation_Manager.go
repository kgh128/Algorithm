/*
1. min heap으로만 풀기
- 시간복잡도: Constructor() - O(nlogn), Reserve()/Unreserve() - O(logn)
- 공강복잡되 O(n)
- 1~n까지의 자리를 min heap에 넣기 (예약 가능한 자리를 저장하는 자료 구조)
- 예약할 때는 min heap의 root를 꺼내서 반환하고, 예약 취소할 때는 min heap에 취소된 자리를 넣기
+) golang에서 heap을 사용하고 싶으면 "container/heap"을 import 해야 하고, golang이 제시하는 인터페이스를 다 구현해야 한다.
*/
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

type SeatManager struct {
    seats IntHeap
}


func Constructor(n int) SeatManager {
    seats := IntHeap{}

    for i := 1; i <= n; i++ {
        heap.Push(&seats, i)
    }

    return SeatManager{seats: seats}
}


func (this *SeatManager) Reserve() int {
    return heap.Pop(&this.seats).(int)
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(&this.seats, seatNumber)
}


/*
2. min heap과 counter 사용
- 시간복잡도: Constructor() - O(1), Reserve()/Unreserve() - counter 사용하면 O(1), min heap 사용하면 O(logn)
- 공강복잡도: O(n)
- min heap은 처음에는 비워두기 (예약 취소된 자리를 위한 자료 구조) -> counter보다 숫자가 작은 자리만 저장한다.
- counter는 현재 예약된 자리 중 가장 숫자가 큰 자리를 저장하고 있다.
- 예약을 했는데 min heap이 비어있다면 남은 자리 중에 counter보다 숫자가 작은 자리는 없다는 의미이므로 counter를 하나 증가시키고 해당 counter의 자리를 예약한다.
- 예약을 취소했는데 해당 자리가 counter와 같다면 가장 숫자가 큰 자리가 취소된 것이므로 counter의 값을 1 감소시킨다.
=> 이때, 1 감소시킨 counter의 값이 꼭 예약중인 자리가 아닐 수도 있다. 하지만 그런 경우 이 값은 이미 min heap에 저장되어 있을 것이므로 프로그램은 원하는 방향으로 동작한다.
- 예약을 취소했는데 해당 자리가 counter보다 작다면 해당 자리를 min heap에 저장한다.
+) IntHeap처럼 사용자 정의 타입도 make() 함수를 이용하여 객체를 만들 수 있다.
*/
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}

type SeatManager struct {
    lastSeat int
    unreservedSeats IntHeap
}


func Constructor(n int) SeatManager {
    return SeatManager{
        lastSeat: 0,
        unreservedSeats: make(IntHeap, 0),
    }
}


func (this *SeatManager) Reserve() int {
    if len(this.unreservedSeats) == 0 {
        this.lastSeat++
        return this.lastSeat
    }
    return heap.Pop(&this.unreservedSeats).(int)
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    if seatNumber == this.lastSeat {
        this.lastSeat--
    } else {
        heap.Push(&this.unreservedSeats, seatNumber)
    }
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
