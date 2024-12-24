package main

import (
	"fmt"
)

// Интерфейс геометрической фигуры
type Shape interface {
	Accept(visitor ShapeVisitor)
}

// Круг
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

// Прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

// Интерфейс посетителя
type ShapeVisitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}

// Посетитель: сохранение в текст
type SaveAsTextVisitor struct{}

func (v *SaveAsTextVisitor) VisitCircle(circle *Circle) {
	fmt.Printf("Circle: Radius = %.2f\n", circle.Radius)
}

func (v *SaveAsTextVisitor) VisitRectangle(rect *Rectangle) {
	fmt.Printf("Rectangle: Width = %.2f, Height = %.2f\n", rect.Width, rect.Height)
}

// Посетитель: сохранение в JSON
type SaveAsJSONVisitor struct{}

func (v *SaveAsJSONVisitor) VisitCircle(circle *Circle) {
	fmt.Printf("{\"type\":\"Circle\",\"radius\":%.2f}\n", circle.Radius)
}

func (v *SaveAsJSONVisitor) VisitRectangle(rect *Rectangle) {
	fmt.Printf("{\"type\":\"Rectangle\",\"width\":%.2f,\"height\":%.2f}\n", rect.Width, rect.Height)
}

// Посетитель: сохранение в XML
type SaveAsXMLVisitor struct{}

func (v *SaveAsXMLVisitor) VisitCircle(circle *Circle) {
	fmt.Printf("<Circle><Radius>%.2f</Radius></Circle>\n", circle.Radius)
}

func (v *SaveAsXMLVisitor) VisitRectangle(rect *Rectangle) {
	fmt.Printf("<Rectangle><Width>%.2f</Width><Height>%.2f</Height></Rectangle>\n", rect.Width, rect.Height)
}

// Посетитель: сохранение в CSV
type SaveAsCSVVisitor struct {
	FilePath string
	Records  [][]string
}

func (v *SaveAsCSVVisitor) VisitCircle(circle *Circle) {
	v.Records = append(v.Records, []string{"Circle", "Radius", fmt.Sprintf("%.2f", circle.Radius)})
}

func (v *SaveAsCSVVisitor) VisitRectangle(rect *Rectangle) {
	v.Records = append(v.Records, []string{"Rectangle", "Width", fmt.Sprintf("%.2f", rect.Width), "Height", fmt.Sprintf("%.2f", rect.Height)})
}

// Метод для вывода CSV в консоль
func (v *SaveAsCSVVisitor) Save() {
	fmt.Println("CSV Data:")
	for _, record := range v.Records {
		fmt.Println(record)
	}
}

func main() {
	// Список фигур
	shapes := []Shape{
		&Circle{Radius: 5.0},
		&Rectangle{Width: 10.0, Height: 20.0},
	}

	// Создаем разные виды посетителей
	textVisitor := &SaveAsTextVisitor{}
	jsonVisitor := &SaveAsJSONVisitor{}
	xmlVisitor := &SaveAsXMLVisitor{}
	csvVisitor := &SaveAsCSVVisitor{}

	// Сохраняем фигуры в текст
	fmt.Println("Сохранение в текст:")
	for _, shape := range shapes {
		shape.Accept(textVisitor)
	}

	// Сохраняем фигуры в JSON
	fmt.Println("\nСохранение в JSON:")
	for _, shape := range shapes {
		shape.Accept(jsonVisitor)
	}

	// Сохраняем фигуры в XML
	fmt.Println("\nСохранение в XML:")
	for _, shape := range shapes {
		shape.Accept(xmlVisitor)
	}

	// Сохраняем фигуры в CSV
	fmt.Println("\nСохранение в CSV:")
	for _, shape := range shapes {
		shape.Accept(csvVisitor)
	}
	csvVisitor.Save()
}
