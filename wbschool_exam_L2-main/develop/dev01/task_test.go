package main

import (
	"testing"
	"time"
)

// TestGetNTPTime_Success тестирует успешное получение времени с сервера NTP.
func TestGetNTPTime_Success(t *testing.T) {
	ntpTime, err := GetNTPTime("pool.ntp.org")
	if err != nil {
		t.Fatalf("Ожидался успешный ответ, но получена ошибка: %v", err)
	}

	// Проверяем, что полученное время не является нулевым значением
	if ntpTime.IsZero() {
		t.Fatalf("Полученное время не должно быть нулевым")
	}

	// Проверяем, что разница с локальным временем не слишком велика
	localTime := time.Now()
	if diff := localTime.Sub(ntpTime).Seconds(); diff > 10 {
		t.Fatalf("Разница между локальным и NTP временем слишком велика: %f секунд", diff)
	}
}

// TestGetNTPTime_Failure тестирует обработку ошибки при некорректном сервере NTP.
func TestGetNTPTime_Failure(t *testing.T) {
	_, err := GetNTPTime("invalid.ntp.server")
	if err == nil {
		t.Fatalf("Ожидалась ошибка при использовании некорректного NTP сервера")
	}
}
