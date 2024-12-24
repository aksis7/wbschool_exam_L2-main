package main

import "fmt"

// Handler — интерфейс обработчика
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

// BaseHandler — базовая структура для обработчиков
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) Handler {
	b.next = handler
	return handler
}

func (b *BaseHandler) Handle(request string) {
	if b.next != nil {
		b.next.Handle(request)
	}
}

// InfoHandler — обработчик для информационных сообщений
type InfoHandler struct {
	BaseHandler
}

func (h *InfoHandler) Handle(request string) {
	if request == "info" {
		fmt.Println("InfoHandler: Обрабатываю информационное сообщение")
	} else {
		fmt.Println("InfoHandler: Передаю дальше")
		h.BaseHandler.Handle(request)
	}
}

// WarningHandler — обработчик для предупреждений
type WarningHandler struct {
	BaseHandler
}

func (h *WarningHandler) Handle(request string) {
	if request == "warning" {
		fmt.Println("WarningHandler: Обрабатываю предупреждение")
	} else {
		fmt.Println("WarningHandler: Передаю дальше")
		h.BaseHandler.Handle(request)
	}
}

// ErrorHandler — обработчик для ошибок
type ErrorHandler struct {
	BaseHandler
}

func (h *ErrorHandler) Handle(request string) {
	if request == "error" {
		fmt.Println("ErrorHandler: Обрабатываю ошибку")
	} else {
		fmt.Println("ErrorHandler: Передаю дальше")
		h.BaseHandler.Handle(request)
	}
}

// Клиентский код
func main() {
	info := &InfoHandler{}
	warning := &WarningHandler{}
	errorHandler := &ErrorHandler{}

	// Формируем цепочку
	info.SetNext(warning).SetNext(errorHandler)

	// Тестируем
	fmt.Println("Отправляем 'info':")
	info.Handle("info")

	fmt.Println("\nОтправляем 'warning':")
	info.Handle("warning")

	fmt.Println("\nОтправляем 'error':")
	info.Handle("error")

	fmt.Println("\nОтправляем 'unknown':")
	info.Handle("unknown")
}
