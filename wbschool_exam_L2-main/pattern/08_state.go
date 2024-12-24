package main

import "fmt"

// State интерфейс, определяющий методы для состояний
type State interface {
	insertCoin()
	pressButton()
	dispense()
}

// Context - автомат по продаже напитков
type VendingMachine struct {
	state State
}

func (vm *VendingMachine) setState(state State) {
	vm.state = state
}

func (vm *VendingMachine) insertCoin() {
	vm.state.insertCoin()
}

func (vm *VendingMachine) pressButton() {
	vm.state.pressButton()
}

func (vm *VendingMachine) dispense() {
	vm.state.dispense()
}

// ConcreteState: NoCoinState
type NoCoinState struct {
	vm *VendingMachine
}

func (s *NoCoinState) insertCoin() {
	fmt.Println("Монета вставлена.")
	s.vm.setState(&HasCoinState{s.vm})
}

func (s *NoCoinState) pressButton() {
	fmt.Println("Вставьте монету сначала.")
}

func (s *NoCoinState) dispense() {
	fmt.Println("Оплатите товар.")
}

// ConcreteState: HasCoinState
type HasCoinState struct {
	vm *VendingMachine
}

func (s *HasCoinState) insertCoin() {
	fmt.Println("Монета уже вставлена.")
}

func (s *HasCoinState) pressButton() {
	fmt.Println("Кнопка нажата. Выдача товара...")
	s.vm.setState(&NoCoinState{s.vm})
}

func (s *HasCoinState) dispense() {
	fmt.Println("Товар выдан.")
}

// Клиентский код
func main() {
	vendingMachine := &VendingMachine{state: &NoCoinState{}}

	vendingMachine.insertCoin()  // Монета вставлена
	vendingMachine.pressButton() // Кнопка нажата
	vendingMachine.dispense()    // Товар выдан

	vendingMachine.pressButton() // Вставьте монету сначала
}
