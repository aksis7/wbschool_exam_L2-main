package main

import "fmt"

// PaymentStrategy — интерфейс стратегии
type PaymentStrategy interface {
	Pay(amount float64)
}

// CreditCardPayment — стратегия для оплаты кредитной картой
type CreditCardPayment struct {
	CardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Оплачено %.2f с использованием кредитной карты: %s\n", amount, c.CardNumber)
}

// PayPalPayment — стратегия для оплаты через PayPal
type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Оплачено %.2f с использованием PayPal: %s\n", amount, p.Email)
}

// BitcoinPayment — стратегия для оплаты биткойнами
type BitcoinPayment struct {
	WalletAddress string
}

func (b *BitcoinPayment) Pay(amount float64) {
	fmt.Printf("Оплачено %.2f с использованием Bitcoin: %s\n", amount, b.WalletAddress)
}

// Context — контекст, использующий стратегию
type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) Pay(amount float64) {
	pc.strategy.Pay(amount)
}

// Main
func main() {
	context := PaymentContext{}

	// Оплата кредитной картой
	context.SetStrategy(&CreditCardPayment{CardNumber: "1234-5678-9012-3456"})
	context.Pay(100.50)

	// Оплата через PayPal
	context.SetStrategy(&PayPalPayment{Email: "user@example.com"})
	context.Pay(200.75)

	// Оплата биткойнами
	context.SetStrategy(&BitcoinPayment{WalletAddress: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"})
	context.Pay(0.05)
}

/*
Необходимость выбора одного из нескольких вариантов поведения во время выполнения.

Когда есть несколько алгоритмов, которые можно применить к одной задаче, и их можно легко заменять.
Избавление от множества условных операторов (if-else или switch-case).

Паттерн помогает сделать код более чистым, заменяя условные блоки на набор классов со схожим интерфейсом.
Расширяемость и поддерживаемость.

Легко добавлять новые стратегии (алгоритмы) без изменения существующего кода.
Изоляция кода алгоритмов.

Стратегии инкапсулируют алгоритмы, что упрощает тестирование и замену.
*/
