type SeatManager struct {
    capacity int
    nextSeat int
    isReserved map[int]bool  
}


func Constructor(n int) SeatManager {
    isReserved := make(map[int]bool)

    for i := 1; i <= n; i++ {
        isReserved[i] = false
    }

    return SeatManager{capacity: n, nextSeat: 1, isReserved: isReserved}
}


func (this *SeatManager) Reserve() int {
    seatNumber := this.nextSeat
    this.isReserved[seatNumber] = true

    for i := seatNumber + 1; i <= this.capacity; i++ {
        if !this.isReserved[i] {
            this.nextSeat = i
            break
        }
    }

    return seatNumber
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    this.isReserved[seatNumber] = false

    if seatNumber < this.nextSeat {
        this.nextSeat = seatNumber
    }
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */
