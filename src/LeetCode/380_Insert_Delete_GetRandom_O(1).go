/*
<map의 key list의 인덱스를 랜덤으로 뽑기>
- 시간복잡도: O(1) -> 문제의 조건!
- 공간복잡도: O(n)

- golang의 rand 함수는 범위를 주면 그 안에서 랜덤으로 숫자르 뽑는 기능만 제공한다.
  => 슬라이스, 배열, 맵 안에 있는 요소 중에서만 랜덤으로 뽑고 싶다면 바로 해당 요소를 랜덤으로 뽑을 수는 없고,
  => 자료구조의 크기를 rand 함수의 범위로 넘겨 인덱스를 랜덤으로 뽑는 방식을 사용해야 한다.
  
- map 같은 경우는 인덱스로 접근하는 것이 아니라 key로 접근하므로 key 슬라이스를 관리하고, key 슬라이스의 인덱스를 랜덤으로 뽑는 방식을 사용한다.
- 이 문제에서는 실질적인 값의 집합은 key 슬라이스이고, map은 해당 key가 존재하는지 확인하는 기능으로 사용한다. (set을 구현)

- key 슬라이스를 얻는 가장 쉬운 방법은 map을 반복문을 통해 순회하면서 한 번에 얻는 것인데, 이는 시간복잡도가 O(n)이어서 사용할 수 없다.
- 그래서 key 슬라이스도 map과 함께 바로 바로 insert와 delete를 수행하면서 관리한다.

- insert 할 때 map으로 확인하고 map에 key-index 쌍, key slice에 key 값을 append한다.

- delete 할 때는 map으로 확인하고, map에서 key-index 쌍을 제거하느 것은 쉬운데 key slice에서 key 값을 제거하는 것이 까다롭다.
  => 중간에 key를 제거하면 해당 key 뒤에 있던 다른 key들의 index가 하나씩 앞으로 당겨지기 때문이다.
  => 그러면 해당 key들은 map에서 key-index 쌍의 index 값을 모두 수정해야 하는데, 이것도 시간복잡도가 O(n)이다.
  
- 그래서 delete 할 때 바로 제거하지 않고, 원하는 key와 슬라이스의 가장 뒤에 존재하는 key의 자리를 서로 바꾼 후 제거한다.
  => 그러면 제거하기 원하는 key는 슬라이스의 가장 뒤에 위치하게 되어 해당 key를 슬라이스에서 지워도 다른 key들의 인덱스가 변화하지 않는다.
  => 자리를 바꾼 후에 원래 가장 뒤에 존재했던 key의 인덱스 값만 map에서 수정해주면 된다.
  => 그러면 O(1)의 시간복잡도로 delete를 수행할 수 있다.
  
- 이렇게 관리하고 있는 key 슬라이스에서 rand 함수를 이용하여 랜덤 인덱스를 뽑고, 그 인덱스로 key 값을 접근하면 랜덤 key 값을 얻을 수 있다.
*/

import "math/rand"

type RandomizedSet struct {
    set map[int]int
    list []int
}


func Constructor() RandomizedSet {
    randomizedSet := RandomizedSet {
        set: make(map[int]int),
        list: []int{},
    }
    return randomizedSet
}


func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.set[val]; !exists {
        this.set[val] = len(this.list)
        this.list = append(this.list, val)
        return true
    }
    return false
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, exists := this.set[val]; exists {
        idx := this.set[val];
        lastIdx := len(this.list) - 1

        this.list[idx], this.list[lastIdx] = this.list[lastIdx], this.list[idx]
        this.set[this.list[idx]] = idx

        this.list = this.list[:lastIdx]
        delete(this.set, val)

        return true
    }
    return false
}


func (this *RandomizedSet) GetRandom() int {
    return this.list[rand.Intn(len(this.list))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
