package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const apiUrl = "https://godville.net/gods/api/"

func main() {

	var (
		fullUrl string
		godName string
		key     string

		err error

		command string
	)

	err = godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	godName = os.Getenv("GODNAME")

	if godName == "" {
		panic("Укажите имя бога в переменной окружения или .env файле")
	}

	key = os.Getenv("KEY")

	if key == "" {
		panic("Укажите API токен в переменной окружения или .env файле")
	}

	fullUrl = apiUrl + godName + "/" + key

	fmt.Printf("%s на связи!\n", godName)

	go trackActivity(fullUrl, 30)

	for {
		command = ""
		_, _ = fmt.Scanf("%s", &command)

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		fmt.Printf("Вы попытались выполнить команду %s\n", command)
	}
}

func trackActivity(url string, rate int) {

	var (
		c    http.Client
		r    *http.Response
		data responseJSON

		lastDiaryEntry string

		lastHealth uint16 = 0
		lastPrana  uint8  = 0
		lastPillar uint16 = 0
		lastTown   string

		whereabouts string

		err error
	)

	for {
		r, err = c.Get(url)

		if err != nil {
			fmt.Printf("Error while making request: %s", err.Error())
		}

		err = json.NewDecoder(r.Body).Decode(&data)

		if err != nil && err != io.EOF {
			fmt.Printf("Error while reading body: %s\n", err.Error())
		}

		if data.Expired == true {
			fmt.Println("Данные устарели! Требуется зайти либо через браузер либо через клиент")
			os.Exit(1)
		}

		if lastDiaryEntry != data.DiaryLast {
			lastDiaryEntry = data.DiaryLast
			fmt.Printf("%s\n", lastDiaryEntry)
		}

		if lastHealth != data.Health || lastPrana != data.Godpower || lastPillar != data.Distance || lastTown != data.TownName {

			if data.TownName == "" {
				whereabouts = fmt.Sprintf("Столб #%d", data.Distance)
			} else {
				whereabouts = fmt.Sprintf("%s (ст. %d)", data.TownName, data.Distance)
			}

			fmt.Printf(
				"[%s] Здоровье: %d/%d; Прана: %d%%\n",
				whereabouts,
				data.Health,
				data.MaxHealth,
				data.Godpower,
			)

			lastHealth = data.Health
			lastPrana = data.Godpower
			lastTown = data.TownName
			lastPillar = data.Distance
		}

		time.Sleep(time.Second * time.Duration(rate))
	}
}
