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


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */