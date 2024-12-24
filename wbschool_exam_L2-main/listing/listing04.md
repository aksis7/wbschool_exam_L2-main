Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```
fatal error: all goroutines are asleep - deadlock!

Ответ:
```
Горутина отправляет данные в канал, но не закрывает его. Главная горутина ждет завершения канала (close(ch)), что приводит к deadlock. Чтобы исправить, необходимо вызвать close(ch) после отправки данных.

```
