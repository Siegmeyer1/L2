package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type House struct {
	rooms  int
	floors int
}

type HouseBuilder struct {
	house *House
}

type IBuilder interface {
	setRooms(int)
	setFloors(int)
	build() *House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{&House{}}
}

func (hb *HouseBuilder) setRooms(rooms int) {
	hb.house.rooms = rooms
}

func (hb *HouseBuilder) setFloors(floors int) {
	hb.house.floors = floors
}

func (hb *HouseBuilder) build() *House {
	return hb.house
}

type Director struct {
	builder *IBuilder
}

func NewDirector(_builder *IBuilder) *Director {
	dir := Director{_builder}
	dir.builder = _builder
	return &dir
}

func (d Director) construct(rooms, floors int) *House {
	(*d.builder).setRooms(rooms)
	(*d.builder).setFloors(floors)
	return (*d.builder).build()
}

func example() {
	var builder IBuilder = NewHouseBuilder()
	director := NewDirector(&builder)
	house := director.construct(1, 1)
	fmt.Println(house.floors, house.rooms)
}
