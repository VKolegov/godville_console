package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"godville/commands"
	"godville/structs"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const apiUrl = "https://godville.net/gods/api/"

var (
	currentData structs.GodvilleData

	eClient      *http.Client
	eCurrentData *structs.ExtendedData

	prevHeroData structs.Hero

	datetimeLayout string
)

func main() {

	var (
		fullUrl string
		godName string
		key     string

		err error
	)

	err = godotenv.Load("settings.cfg")

	if err != nil {
		panic(err)
	}

	godName = os.Getenv("GODNAME")

	if godName == "" {
		panic("Укажите имя бога в переменной окружения или settings.cfg файле")
	}

	key = os.Getenv("KEY")

	if key == "" {
		panic("Укажите API токен в переменной окружения или settings.cfg файле")
	}

	fullUrl = apiUrl + godName + "/" + key

	password := os.Getenv("PASSWORD")

	if password != "" {

		err = login(godName, password)

		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Не указан пароль. Режим ограниченного функционала")
	}

	datetimeLayout = os.Getenv("DATETIME_FORMAT")

	if eClient != nil {
		go trackExtended(10)
	} else {
		go trackBasic(fullUrl, 30)
	}

	processCommands()
}

func greetings() {

	var (
		godvilleTimeLayout = "2006-01-02T15:04:05-07:00" // почти ISO8601. ISO8601:"-07:00" godville:"-07:00"
		dateFormat         = "2006/01/02"
		timezone           = "Asia/Krasnoyarsk" // TODO: detect from system
		loc                *time.Location

		err error
	)

	fmt.Printf("%s на связи!\n", currentData.Godname)

	loc, err = time.LoadLocation(timezone)

	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	if currentData.TempleCompletedAt != "" {

		templeDate, _ := time.ParseInLocation(
			godvilleTimeLayout,
			currentData.TempleCompletedAt,
			loc,
		)

		fmt.Printf(
			"Храм достроен %s, поздравляем!\n",
			templeDate.Format(dateFormat),
		)
	}

	if currentData.ArkCompletedAt != "" {
		templeDate, _ := time.ParseInLocation(
			godvilleTimeLayout,
			currentData.ArkCompletedAt,
			loc,
		)

		fmt.Printf(
			"Ковчег достроен %s, поздравляем!\n",
			templeDate.Format(dateFormat),
		)
	}

	commandList()
}

func processCommands() {
	var (

		r *bufio.Reader

		input string
		words []string

		command string
		parameters []string

		err error
	)

	r = bufio.NewReader(os.Stdin)
	parameters = make([]string, 8)

	for {
		command = ""
		input, _ = r.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		input = strings.ToLower(input)

		words = strings.Split(input, " ")

		command = words[0]

		if len(words) > 1 {
			parameters = words[1:]
		} else {
			parameters = make([]string, 4)
		}

		switch command {
		case "герой":
			if eCurrentData == nil {
				commands.Hero(currentData)
			} else {
				commands.HeroExtended(eCurrentData)
			}
		case "инв", "инвентарь":
			if eCurrentData == nil {
				commands.Inventory(currentData)
			} else {
				commands.InventoryExtended(eCurrentData)
			}
		case "исп", "п", "предмет":
			if eCurrentData == nil {
				fmt.Println("Недоступно в ограниченной версии")
			} else {
				var id int
				id, err = strconv.Atoi(parameters[0])

				if err != nil {
					fmt.Print("Не удалось распознанть номер предмета")
					continue
				}

				commands.UseItem(id, eCurrentData, eClient)
			}
		case "снар", "снаряжение":
			if eCurrentData == nil {
				fmt.Println("Недоступно в ограниченной версии")
			} else {
				commands.Equipment(eCurrentData)
			}
		case "квест":
			if eCurrentData == nil {
				commands.QuestStatus(currentData)
			} else {
				commands.QuestStatusExtended(eCurrentData.Hero)
			}
		case "зло":
			if eCurrentData == nil {
				fmt.Println("Недоступно в ограниченной версии")
			} else {
				commands.MakeInfluence("punish", eCurrentData, eClient)
			}
		case "добро":
			if eCurrentData == nil {
				fmt.Println("Недоступно в ограниченной версии")
			} else {
				commands.MakeInfluence("encourage", eCurrentData, eClient)
			}
		case "оживить", "воскресить", "воскр":
			if eCurrentData == nil {
				fmt.Println("Недоступно в ограниченной версии")
			} else {
				commands.ResurrectHero(eClient, eCurrentData)
			}
		case "бог", "я":
			if eCurrentData == nil {
				commands.PrintGodInfo(currentData, true)
			} else {
				commands.PrintGodInfo(eCurrentData.Hero, true)
			}
		case "команды":
			commandList()
		case "выход":
			os.Exit(0)

		default:
			fmt.Printf("Вы попытались выполнить команду %s\n", command)
		}
	}
}

func commandList() {
	fmt.Println()
	fmt.Println("Команды:")
	fmt.Println("	'герой' 							- вывести информацию о герое")
	fmt.Println("	'инвентарь' или 'инв' 				- вывести информацию об инвентаре героя")
	fmt.Println("	'предмет', 'исп' или 'п' 			- активировать предмет в инвентаре")
	fmt.Println("	'снаряжение' или 'снар' 			- вывести информацию о снаряжении героя (недоступно в огранич. версии)")
	fmt.Println("	'бог' или 'я' 						- вывести информацию об себе (божестве)")
	fmt.Println("	'квест' 							- вывести информацию о текущем задании")
	fmt.Println("	'добро' 							- сделать добро (недоступно в огранич. версии)")
	fmt.Println("	'зло' 								- сделать зло (недоступно в огранич. версии)")
	fmt.Println("	'оживить', 'воскресить', 'воскр'	- воскресить героя (недоступно в огранич. версии)")
	fmt.Println("	'команды' 							- вывести список команд")
	fmt.Println("	'выход' 							- закрыть программу")
	fmt.Println()
}
