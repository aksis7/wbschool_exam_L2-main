package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// GetNTPTime получает текущее точное время с NTP сервера.
func GetNTPTime(server string) (time.Time, error) {
	return ntp.Time(server)
}

func main() {
	// Получаем точное время с NTP сервера
	ntpTime, err := GetNTPTime("pool.ntp.org")
	if err != nil {
		// Используем fmt.Fprintf для вывода ошибки в STDERR
		fmt.Fprintf(os.Stderr, "Ошибка получения точного времени: %v\n", err)
		os.Exit(1)
	}

	// Локальное системное время
	localTime := time.Now()

	// Выводим системное и точное время
	fmt.Println("Локальное системное время:", localTime.Format(time.RFC1123))
	fmt.Println("Точное NTP время:", ntpTime.Format(time.RFC1123))
}
