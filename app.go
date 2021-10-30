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

var (
	currentData responseJSON
)

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

		if currentData.Expired == true {
			fmt.Println("Данные устарели! Требуется зайти либо через браузер либо через клиент")
			os.Exit(1)
		}

		trackBasicData()

		if currentData.TempleCompletedAt == "" {
			trackBricks()
		}

		if currentData.ArkCompletedAt == "" {
			trackWood()
		}

		time.Sleep(time.Second * time.Duration(rate))
	}
}

var (
	lastDiaryEntry string

	lastHealth uint16 = 0
	lastPrana  uint8  = 0
	lastPillar uint16 = 0
	lastTown   string

	lastBrickCnt int16 = -1
	lastWoodCnt  int32 = -1
)

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

func trackBasicData() {

	var (
		whereabouts string
	)

	if lastDiaryEntry != currentData.DiaryLast {
		lastDiaryEntry = currentData.DiaryLast
		fmt.Printf("[Дневник] %s\n", lastDiaryEntry)
	}

	if lastHealth != currentData.Health || lastPrana != currentData.Godpower || lastPillar != currentData.Distance || lastTown != currentData.TownName {

		if currentData.TownName == "" {
			whereabouts = fmt.Sprintf("Столб #%d", currentData.Distance)
		} else {
			whereabouts = fmt.Sprintf("%s (ст. %d)", currentData.TownName, currentData.Distance)
		}

		fmt.Printf(
			"[%s] Здоровье: %d/%d; Прана: %d%%",
			whereabouts,
			currentData.Health,
			currentData.MaxHealth,
			currentData.Godpower,
		)

		if currentData.BricksCnt < 1000 {
			fmt.Printf("; Золотых кирпичей: %d/1000", currentData.BricksCnt)
		}

		if currentData.WoodCnt < 1000 {
			fmt.Printf("; Дерева: %d/1000", currentData.WoodCnt)
		}

		fmt.Print("\n")

		lastHealth = currentData.Health
		lastPrana = currentData.Godpower
		lastTown = currentData.TownName
		lastPillar = currentData.Distance
	}

}

func trackBricks() {

	var bCnt, diff int16
	bCnt = int16(currentData.BricksCnt)

	if lastWoodCnt == -1 {
		lastBrickCnt = bCnt
	} else if lastBrickCnt != bCnt {
		diff = bCnt - lastBrickCnt
		lastBrickCnt = bCnt
		fmt.Printf("%s получил %d кирпичей!\n", currentData.Name, diff)
	}

	if currentData.BricksCnt == 1000 {
		fmt.Printf("%s собрал все кирпичи для постройки храма!\n", currentData.Name)
	}
}

func trackWood() {

	var wCnt, diff int32
	wCnt = int32(currentData.WoodCnt)

	if lastWoodCnt == -1 {
		lastWoodCnt = wCnt
	} else if lastWoodCnt != wCnt {
		diff = wCnt - lastWoodCnt
		lastWoodCnt = wCnt
		fmt.Printf("Герой получил %d досок для ковчега!\n", diff)
	}

	if currentData.WoodCnt == 1000 {
		fmt.Printf("%s собрал дерево для постройки ковчега!\n", currentData.Name)
	}
}
