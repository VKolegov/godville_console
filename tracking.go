package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"godville/commands"
	"io"
	"net/http"
	"time"
)

var (
	lastDiaryEntry       string
	lastNewsFromTheField string

	lastHealth uint16 = 0
	lastPrana  uint8  = 0
	lastGold   string
	lastPillar uint16 = 0
	lastTown   string

	lastBrickCnt int16 = -1
	lastWoodCnt  int32 = -1

	lastSavings       int32 = -1
	lastSavingsString string

	//lastMonsterName string
	lastMonsterProgress uint16
)

func trackBasic(url string, rate int) {

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

func trackGodData() {

	if lastPrana != currentData.Godpower ||
		(lastBrickCnt >= 0 && lastBrickCnt != int16(currentData.BricksCnt)) ||
		(lastWoodCnt >= 0 && lastWoodCnt != int32(currentData.WoodCnt)) ||
		lastSavingsString != currentData.Savings {

		commands.GodInfo(currentData)

		lastPrana = currentData.Godpower
		lastSavingsString = currentData.Savings
	}
}

func trackHeroData() {
	var (
		whereabouts string
	)

	if lastDiaryEntry != currentData.DiaryLast {
		lastDiaryEntry = currentData.DiaryLast
		fmt.Printf("[%s][Дневник] %s\n", currentData.Name, lastDiaryEntry)
	}

	if lastHealth != currentData.Health ||
		lastPillar != currentData.Distance ||
		lastTown != currentData.TownName ||
		lastGold != currentData.GoldApprox {

		if currentData.TownName == "" {
			whereabouts = fmt.Sprintf("Столб #%d", currentData.Distance)
		} else {
			whereabouts = fmt.Sprintf("%s (ст. %d)", currentData.TownName, currentData.Distance)
		}

		fmt.Printf(
			"[%s][%s] Здоровье: %d/%d; Золота: %s\n",
			currentData.Name,
			whereabouts,
			currentData.Health,
			currentData.MaxHealth,
			currentData.GoldApprox,
		)

		lastHealth = currentData.Health
		lastTown = currentData.TownName
		lastPillar = currentData.Distance
		lastGold = currentData.GoldApprox
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

func trackSavings() {

	savings, err := parseSavings(currentData.Savings)

	if err != nil {
		fmt.Println(err)
		return
	}

	if lastSavings == -1 {
		lastSavings = savings
	} else if lastSavings != savings {

		diff := savings - lastSavings

		fmt.Printf("Герой отложил %d тысяч!", diff)
	}
}

func parseSavings(savingsString string) (int32, error) {
	var (
		savings = 0
		a       string // a = тысяч
	)
	_, err := fmt.Sscanf(savingsString, "%d %s", &savings, &a)

	if err != nil {

		return -1, errors.New(
			fmt.Sprintf(
				"Ошибка при парсинге сбережений \"%s\": %s\n", savingsString, err.Error(),
			),
		)
	}

	return int32(savings), nil
}
