package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"time"
)

const apiUrl = "https://godville.net/gods/api/"

func main() {

	var (
		fullUrl string
		godName string
		key     string

		r            *http.Response
		responseJSON responseJSON

		lastHealth     uint16 = 0
		lastPrana      uint8  = 0
		lastDiaryEntry string = ""

		err error
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

	c := http.Client{}

	for {
		r, err = c.Get(fullUrl)

		if err != nil {
			fmt.Printf("Error while making request: %s", err.Error())
		}

		err = json.NewDecoder(r.Body).Decode(&responseJSON)

		if err != nil && err != io.EOF {
			fmt.Printf("Error while reading body: %s\n", err.Error())
		}

		if lastDiaryEntry != responseJSON.DiaryLast {
			lastDiaryEntry = responseJSON.DiaryLast
			fmt.Printf("%s\n", lastDiaryEntry)
		}

		if lastHealth != responseJSON.Health || lastPrana != responseJSON.Godpower {
			fmt.Printf(
				"Здоровье: %d/%d; Прана: %d%%\n",
				responseJSON.Health,
				responseJSON.MaxHealth,
				responseJSON.Godpower,
			)

			lastHealth = responseJSON.Health
			lastPrana = responseJSON.Godpower
		}

		time.Sleep(time.Second * 30)
	}

}
