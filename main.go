package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Описание игрока
type Player struct {
	Inventory []string // Инвентарь игрока
	Location  string   // Текущая локация игрока
}

// Локация в игре
type Location struct {
	Name        string   // Название локации
	Description string   // Описание локации
	Items       []string // Предметы в локации
	Exits       []string // Доступные направления для перемещения
}

// Глобальные переменные
var (
	player    *Player
	locations map[string]*Location
)

// Инициализирует игровой мир
func initGame() {
	player = &Player{
		Inventory: make([]string, 0),
		Location:  "кухня", // Начальная локация игрока
	}

	locations = make(map[string]*Location)
	locations["кухня"] = &Location{
		Name:        "кухня",
		Description: "кухня, ничего интересного. можно пройти - коридор",
		Items:       []string{},
		Exits:       []string{"коридор"},
	}
	locations["коридор"] = &Location{
		Name:        "коридор",
		Description: "ничего интересного. можно пройти - кухня, комната, улица",
		Items:       []string{},
		Exits:       []string{"кухня", "комната", "улица"},
	}
	locations["комната"] = &Location{
		Name:        "комната",
		Description: "ты в своей комнате. можно пройти - коридор",
		Items:       []string{"ключи", "конспекты", "рюкзак"},
		Exits:       []string{"коридор"},
	}
	locations["улица"] = &Location{
		Name:        "улица",
		Description: "на улице весна. можно пройти - домой",
		Items:       []string{},
		Exits:       []string{"коридор"},
	}

}

func main() {
	initGame()
	fmt.Println("Добро пожаловать в текстовую игру!")
	fmt.Println("Доступные команды: осмотреться, идти <направление>, взять <предмет>")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка при чтении ввода:", err)
			continue
		}
		fmt.Println(handleCommand(input))
	}

}

// Обрабатываем команды игрока
type GameWorld struct{}

func handleCommand(input string) string {
	command := strings.TrimSpace(input)
	parts := strings.Split(command, " ")
	command = parts[0]
	params := parts[1:]

	gw := &GameWorld{}

	switch command {
	case "осмотреться":
		return gw.LookAround()
	case "идти":
		return gw.GoTo(params)
	case "взять":
		return gw.Take(params)
	default:
		return "неизвестная команда"
	}
}

func (gw *GameWorld) LookAround() string {
	if player.Location == "комната" {
		var items string
		for _, v := range locations[player.Location].Items {
			items += v + ", "
		}
		if len(items) > 0 {
			// Удаляем лишнюю запятую и пробел в конце строки
			items = items[:len(items)-2]
		}
		if len(items) == 0 {
			return fmt.Sprintf("пустая комната. можно пройти - коридор")
		}
		return fmt.Sprintf("на столе: %s. можно пройти - коридор", items)
	}
	if player.Location == "кухня" {
		return fmt.Sprintf("ты находишься на кухне, надо собрать рюкзак и идти в универ. можно пройти - коридор")
	}
	return locations[player.Location].Description
}

func (gw *GameWorld) GoTo(params []string) string {
	if len(params) == 0 {
		return "Некорректное направление"
	}
	nextLocation := params[0]
	for _, exit := range locations[player.Location].Exits {
		if nextLocation == exit {
			player.Location = nextLocation
			return fmt.Sprintf("%s", locations[player.Location].Description)
		}
	}
	return fmt.Sprintf("нет пути в %s", params[0])
}

func (gw *GameWorld) Take(params []string) string {
	if len(params) == 0 {
		return "Некорректное название предмета"
	}
	item := params[0]
	for index, currentItem := range locations[player.Location].Items {
		if item == currentItem {
			player.Inventory = append(player.Inventory, item)
			locations[player.Location].Items = append(locations[player.Location].Items[:index], locations[player.Location].Items[index+1:]...)
			return fmt.Sprintf("предмет добавлен в инвентарь: %s", item)
		}
	}
	return "нет такого"
}
