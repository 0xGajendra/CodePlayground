package main

import (
	"fmt"
	"math/rand"
)

type Fighter interface{
	Attack() int
	GetName() string
	GetHP() int
	TakeDamage(int)
}

type Hero struct{
	name string
	health int
}

type Monster struct{
	name string
	health int
}

func (h *Hero) GetName() string {
	return h.name
}

func (h *Hero) GetHP() int {
	return h.health
}

func (h *Hero) Attack() int {
	return rand.Int()%100
}

func (h *Hero) TakeDamage(n int) {
	h.health-=n
}

func (m *Monster) GetName() string {
	return m.name
}

func (m *Monster) GetHP() int {
	return m.health
}

func (m *Monster) Attack() int {
	return rand.Int()%150
}

func (m *Monster) TakeDamage(n int) {
	m.health-=n
}

func Battle( f1 Fighter, f2 Fighter){
	fmt.Printf("Starting a fight: %v vs %v", f1.GetName(), f2.GetName())

	for f1.GetHP()>0 && f2.GetHP()>0 {
		fmt.Printf("%v attacked %v\n\n", f1.GetName(), f2.GetName())
		hp1 := f1.Attack()
		fmt.Printf("%v took a damage of %v from %v\n\n", f1.GetName(),hp1, f2.GetName())
		f2.TakeDamage(hp1)
		if(f2.GetHP()<=0){
			fmt.Printf("%v lost the battle and %v won\n\n", f2.GetName(), f1.GetName())
			break;
		}


		fmt.Printf("%v attacked %v\n\n", f2.GetName(), f1.GetName())
		hp2 := f2.Attack()
		fmt.Printf("%v took a damage of %v from %v\n\n", f2.GetName(),hp2, f1.GetName())

		f1.TakeDamage(hp2)
		if(f1.GetHP()<=0){
			fmt.Printf("%v lost the battle and %v won\n]n", f1.GetName(), f2.GetName())
			break
		}

	}
}


func main(){
	hero:=Hero{
		name: "tenma",
		health: 1000,
	}

	monster:=Monster{
		name: "johan libert",
		health: 677,
	}

	Battle(&hero, &monster)

	fmt.Printf("Fight ended\n\n")

}