package models

import (
	"container/heap"
	"fmt"
)

type Parking struct {
	Available        *IntMinHeap
	Occupied         map[int]Car
	slotByPlatNumber map[string]int
}

func NewParking(n int) *Parking {
	h := &IntMinHeap{}
	for i := 1; i <= n; i++ {
		*h = append(*h, i)
	}
	heap.Init(h)

	return &Parking{
		Available:        h,
		Occupied:         make(map[int]Car),
		slotByPlatNumber: make(map[string]int),
	}
}

func (p *Parking) Park(platNumber string) (int, error) {
	if p.Available.Len() == 0 {
		return 0, fmt.Errorf("sorry, parking lot full")
	}

	slot := heap.Pop(p.Available).(int)

	car := Car{PlateNumber: platNumber, Colour: "red"}
	p.Occupied[slot] = car
	p.slotByPlatNumber[platNumber] = slot

	return slot, nil
}

func (p *Parking) Leave(platNumber string) (int, error) {
	slot, ok := p.slotByPlatNumber[platNumber]
	if !ok {
		return 0, fmt.Errorf("car not found")
	}

	delete(p.slotByPlatNumber, platNumber)
	delete(p.Occupied, slot)

	heap.Push(p.Available, slot)

	return slot, nil
}
