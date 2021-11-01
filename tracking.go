package main

import (
	"encoding/json"
	"fmt"
	"godville/displaying"
	"io"
	"net/http"
	"time"
)

var (
	lastDiaryEntry       string
	lastNewsFromTheField string

	lastHealth  uint16 = 0
	lastPrana   uint8  = 0
	lastGoldStr string
	lastGold    int
	lastInvNum  uint16
	lastPillar  uint16 = 0
	lastTown    string

	lastBrickCnt int16 = -1
	lastWoodCnt  int32 = -1

	lastSavingsString string

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
			trackSavings(currentData, prevHeroData)
		}

		time.Sleep(time.Second * time.Duration(rate))
	}
}

func trackGodData() {

	if lastPrana != currentData.Godpower ||
		(lastBrickCnt >= 0 && lastBrickCnt != int16(currentData.BricksCnt)) ||
		(lastWoodCnt >= 0 && lastWoodCnt != int32(currentData.WoodCnt)) ||
		lastSavingsString != currentData.Savings {

		displaying.PrintGodStatus(currentData, false, prevHeroData, datetimeLayout)

		lastPrana = currentData.Godpower
		lastSavingsString = currentData.Savings
	}
}

func trackHeroData() {

	if lastDiaryEntry != currentData.DiaryLast {
		lastDiaryEntry = currentData.DiaryLast
		fmt.Printf("[%s][Дневник] %s\n", currentData.Name, lastDiaryEntry)
	}

	if lastHealth != currentData.Health ||
		lastPillar != currentData.Distance ||
		lastTown != currentData.TownName ||
		lastGoldStr != currentData.GoldApprox {

		displaying.PrintHeroStatus(currentData, prevHeroData, datetimeLayout)

		lastHealth = currentData.Health
		lastTown = currentData.TownName
		lastPillar = currentData.Distance
		lastGoldStr = currentData.GoldApprox
		lastInvNum = currentData.InventoryNum
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
