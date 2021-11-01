package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"godville/structs"
	"net/http"
	"os"
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
