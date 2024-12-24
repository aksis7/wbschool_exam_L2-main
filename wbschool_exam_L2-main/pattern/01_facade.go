package main

import "fmt"

// Подсистема CPU
type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU: Freezing")
}

func (c *CPU) Jump(position int) {
	fmt.Printf("CPU: Jumping to position %d\n", position)
}

func (c *CPU) Execute() {
	fmt.Println("CPU: Executing")
}

// Подсистема HardDrive
type HardDrive struct{}

func (h *HardDrive) Read(lba int, size int) string {
	fmt.Printf("HardDrive: Reading data at LBA %d with size %d\n", lba, size)
	return "data"
}

// Подсистема Memory
type Memory struct{}

func (m *Memory) Load(position int, data string) {
	fmt.Printf("Memory: Loading data '%s' to position %d\n", data, position)
}

// Фасад
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *ComputerFacade) Start() {
	fmt.Println("ComputerFacade: Starting computer")
	c.cpu.Freeze()
	data := c.hardDrive.Read(0, 1024)
	c.memory.Load(0, data)
	c.cpu.Jump(0)
	c.cpu.Execute()
}

// Главная функция
func main() {
	computer := NewComputerFacade()
	computer.Start()
}

/*
Применимость
1 Подсистемы сложны для понимания или их интерфейсы неудобны для прямого использования.
2 Требуется унифицировать и упростить доступ к функциональности подсистемы.
3 Необходимо скрыть детали реализации подсистемы от клиентов.
Плюсы
1 Снижение сложности:
	Клиентский код взаимодействует с упрощённым интерфейсом.
2 Инкапсуляция:
	Реализация подсистем остаётся скрытой.
3 Снижение связности:
	 Клиент знает только о фасаде, а не о конкретных классах подсистемы.
Минусы
1 Может стать "божественным объектом", если в фасад добавляется слишком много логики.
2 Упрощённый интерфейс может ограничить доступ к более специализированной функциональности подсистемы.
*/
