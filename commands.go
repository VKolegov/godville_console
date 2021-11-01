package main

import (
	"bufio"
	"fmt"
	"godville/commands"
	"os"
	"strconv"
	"strings"
)

func commandList() {
	fmt.Println()
	fmt.Println("Команды:")
	fmt.Println("	'герой' 							- вывести информацию о герое")
	fmt.Println("	'инвентарь', 'инв' 					- вывести информацию об инвентаре героя")
	fmt.Println("	'предмет', 'исп', 'п' #номер		- активировать предмет в инвентаре")
	fmt.Println("	'снаряжение', 'снар' 				- вывести информацию о снаряжении героя (недоступно в огранич. версии)")
	fmt.Println("	'квест' 							- вывести информацию о текущем задании")
	fmt.Println("	'бог' или 'я' 						- вывести информацию об себе (божестве)")
	fmt.Println("	'глас', 'г' #фраза 					- глас (недоступно в огранич. версии)")
	fmt.Println("	'добро' 							- сделать добро (недоступно в огранич. версии)")
	fmt.Println("	'зло' 								- сделать зло (недоступно в огранич. версии)")
	fmt.Println("	'оживить', 'воскресить', 'воскр'	- воскресить героя (недоступно в огранич. версии)")
	fmt.Println("	'команды' 							- вывести список команд")
	fmt.Println("	'выход' 							- закрыть программу")
	fmt.Println()
}

func processCommands() {
	var (
		r *bufio.Reader

		input string
		words []string

		command    string
		parameters []string

		err error
	)

	r = bufio.NewReader(os.Stdin)
	parameters = make([]string, 0, 4)

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
			parameters = make([]string, 0, 4)
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
		case "глас", "г":
			if eCurrentData == nil {
				fmt.Println("Недоступно в ограниченной версии")
			} else {

				if len(parameters) == 0 {
					fmt.Println("А что сказать-то хотел?")
					continue
				}

				ph := parameters[0]

				if len(parameters) > 1 {
					ph = strings.Join(parameters, " ")
				}

				commands.GodPhrase(ph, eClient, eCurrentData)
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
