package main

import "fmt"

// Computer — объект, который мы будем строить
type Computer struct {
	CPU string
	RAM string
	GPU string
}

// Метод для вывода информации о компьютере
func (c Computer) Specs() string {
	return fmt.Sprintf("CPU: %s, RAM: %s, GPU: %s", c.CPU, c.RAM, c.GPU)
}

// Builder — интерфейс строителя
type Builder interface {
	SetCPU()
	SetRAM()
	SetGPU()
	GetComputer() Computer
}

// GamingComputerBuilder — строитель для игрового компьютера
type GamingComputerBuilder struct {
	computer Computer
}

func (b *GamingComputerBuilder) SetCPU() {
	b.computer.CPU = "Intel Core i9"
}

func (b *GamingComputerBuilder) SetRAM() {
	b.computer.RAM = "32GB"
}

func (b *GamingComputerBuilder) SetGPU() {
	b.computer.GPU = "NVIDIA RTX 4090"
}

func (b *GamingComputerBuilder) GetComputer() Computer {
	return b.computer
}

// OfficeComputerBuilder — строитель для офисного компьютера
type OfficeComputerBuilder struct {
	computer Computer
}

func (b *OfficeComputerBuilder) SetCPU() {
	b.computer.CPU = "Intel Core i5"
}

func (b *OfficeComputerBuilder) SetRAM() {
	b.computer.RAM = "16GB"
}

func (b *OfficeComputerBuilder) SetGPU() {
	b.computer.GPU = "Integrated Graphics"
}

func (b *OfficeComputerBuilder) GetComputer() Computer {
	return b.computer
}

// Director — управляет процессом строительства
type Director struct {
	builder Builder
}

func (d *Director) SetBuilder(b Builder) {
	d.builder = b
}

func (d *Director) BuildComputer() Computer {
	d.builder.SetCPU()
	d.builder.SetRAM()
	d.builder.SetGPU()
	return d.builder.GetComputer()
}

func main() {
	// Создаем директора
	director := &Director{}

	// Создаем игровой компьютер
	gamingBuilder := &GamingComputerBuilder{}
	director.SetBuilder(gamingBuilder)
	gamingComputer := director.BuildComputer()
	fmt.Println("Gaming Computer:", gamingComputer.Specs())

	// Создаем офисный компьютер
	officeBuilder := &OfficeComputerBuilder{}
	director.SetBuilder(officeBuilder)
	officeComputer := director.BuildComputer()
	fmt.Println("Office Computer:", officeComputer.Specs())
}
