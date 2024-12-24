package main

import "fmt"

// Room — абстрактный интерфейс комнаты
type Room interface {
	Connect(room Room)
}

// MagicRoom — магическая комната
type MagicRoom struct{}

func (m *MagicRoom) Connect(room Room) {
	fmt.Println("MagicRoom connected to another room.")
}

// OrdinaryRoom — обычная комната
type OrdinaryRoom struct{}

func (o *OrdinaryRoom) Connect(room Room) {
	fmt.Println("OrdinaryRoom connected to another room.")
}

// MazeGame — абстрактная структура для игры-лабиринта
type MazeGame interface {
	MakeRoom() Room
}

// BaseMazeGame — базовая реализация для MazeGame
type BaseMazeGame struct {
	Rooms []Room
}

// NewBaseMazeGame — конструктор для базовой игры
func NewBaseMazeGame(game MazeGame) *BaseMazeGame {
	room1 := game.MakeRoom()
	room2 := game.MakeRoom()
	room1.Connect(room2)

	return &BaseMazeGame{
		Rooms: []Room{room1, room2},
	}
}

// MagicMazeGame — конкретная игра с магическими комнатами
type MagicMazeGame struct{}

func (m *MagicMazeGame) MakeRoom() Room {
	return &MagicRoom{}
}

// OrdinaryMazeGame — конкретная игра с обычными комнатами
type OrdinaryMazeGame struct{}

func (o *OrdinaryMazeGame) MakeRoom() Room {
	return &OrdinaryRoom{}
}

// main — клиентский код
func main() {
	fmt.Println("Creating a Magic Maze Game:")
	magicGame := NewBaseMazeGame(&MagicMazeGame{})
	for _, room := range magicGame.Rooms {
		fmt.Printf("Room type: %T\n", room)
	}

	fmt.Println("\nCreating an Ordinary Maze Game:")
	ordinaryGame := NewBaseMazeGame(&OrdinaryMazeGame{})
	for _, room := range ordinaryGame.Rooms {
		fmt.Printf("Room type: %T\n", room)
	}
}
