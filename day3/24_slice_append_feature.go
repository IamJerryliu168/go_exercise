package main

import "fmt"


// append: double to extend the capacity
func main() {

	s1 := make([]int, 0)
	oldCapacaity := cap(s1)

	for i:=0;i<20;i++{
		s1 = append(s1, i)
		if newCapacity := cap(s1); newCapacity > oldCapacaity {
			fmt.Printf("oldCapacaity:%d >>>>>>>>> newCapacity:%d \n", oldCapacaity, newCapacity)
			oldCapacaity = newCapacity
		}
	}

}