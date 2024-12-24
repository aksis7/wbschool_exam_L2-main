package main

import "fmt"

// 1. Command (интерфейс команды)
type Command interface {
	Execute()
}

// 2. Receiver (получатель команды)
type Light struct{}

func (l *Light) On() {
	fmt.Println("Лампочка включена")
}

func (l *Light) Off() {
	fmt.Println("Лампочка выключена")
}

// 3. ConcreteCommand (конкретная команда для включения света)
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// ConcreteCommand (конкретная команда для выключения света)
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// 4. Invoker (инициатор)
type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

// 5. Client (создаём и связываем команды)
func main() {
	light := &Light{}

	turnOn := &LightOnCommand{light: light}
	turnOff := &LightOffCommand{light: light}

	buttonOn := &Button{command: turnOn}
	buttonOff := &Button{command: turnOff}

	buttonOn.Press()  // Вывод: Лампочка включена
	buttonOff.Press() // Вывод: Лампочка выключена
}
