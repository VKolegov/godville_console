package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"godville/commands"
	"godville/structs"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const apiUrl = "https://godville.net/gods/api/"

var (
	currentData structs.GodvilleData
)

func main() {

	var (
		fullUrl string
		godName string
		key     string

		err error

		command string
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

	go trackActivity(fullUrl, 30)

	for {
		command = ""
		_, _ = fmt.Scanf("%s", &command)

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		command = strings.ToLower(command)

		switch command {
		case "квест":
			commands.QuestStatus(currentData)
		case "инв":
			commands.Inventory(currentData)
		case "инвентарь":
			commands.Inventory(currentData)
		case "герой":
			commands.Hero(currentData)
		default:
			fmt.Printf("Вы попытались выполнить команду %s\n", command)
		}

	}
}

func trackActivity(url string, rate int) {

	var (
		c http.Client
		r *http.Response

		initialRequest = true

		err error
	)

	for {
		r, err = c.Get(url)

		if err != nil {
			fmt.Printf("Error while making request: %s", err.Error())
		}

		err = json.NewDecoder(r.Body).Decode(&currentData)

		if err != nil && err != io.EOF {
			fmt.Printf("Error while reading body: %s\n", err.Error())
		}

		if initialRequest {

			greetings()

			initialRequest = false
		}

		// Превышена частота запросов к серверу
		if currentData.Name != "" && currentData.Godname == "" {
			fmt.Println(currentData.Name)
			time.Sleep(time.Minute)
			continue
		}

		if currentData.Expired == true {
			fmt.Println("Данные устарели! Требуется зайти либо через браузер либо через клиент")
			time.Sleep(time.Minute)
			continue
		}

		trackGodData()
		trackHeroData()

		if currentData.TempleCompletedAt == "" {
			trackBricks()
		}

		if currentData.ArkCompletedAt == "" {
			trackWood()
		}

		if currentData.Savings != "" {
			trackSavings()
		}

		time.Sleep(time.Second * time.Duration(rate))
	}
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
}
