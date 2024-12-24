package main

import (
	"fmt"
	"time"
)

// or принимает несколько каналов и возвращает один канал,
// который закрывается, как только любой из входных каналов закрывается.
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если каналов нет, возвращаем закрытый канал.
		closedChan := make(chan interface{})
		close(closedChan)
		return closedChan
	case 1:
		// Если только один канал, возвращаем его напрямую.
		return channels[0]
	case 2:
		// Если два канала, используем select.
		one, two := channels[0], channels[1]
		result := make(chan interface{})
		go func() {
			defer close(result)
			select {
			case <-one:
			case <-two:
			}
		}()
		return result
	default:
		// Если больше двух каналов, разбиваем массив на две части и объединяем их рекурсивно.
		mid := len(channels) / 2
		return or(
			or(channels[:mid]...),
			or(channels[mid:]...),
		)
	}
}

// Пример функции, которая создает канал, закрывающийся через определённое время.
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
