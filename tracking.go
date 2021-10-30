package main

import (
	"errors"
	"fmt"
)

var (
	lastDiaryEntry string

	lastHealth uint16 = 0
	lastPrana  uint8  = 0
	lastGold   string
	lastPillar uint16 = 0
	lastTown   string

	lastBrickCnt int16 = -1
	lastWoodCnt  int32 = -1

	lastSavings       int32 = -1
	lastSavingsString string
)

func trackGodData() {

	if lastPrana != currentData.Godpower ||
		lastBrickCnt != int16(currentData.BricksCnt) ||
		lastWoodCnt != int32(currentData.WoodCnt) ||
		lastSavingsString != currentData.Savings {

		fmt.Printf("[%s] Прана: %d%%", currentData.Godname, currentData.Godpower)

		if currentData.Savings != "" {
			fmt.Printf("; Сбережений: %s", currentData.Savings)
		}

		if currentData.BricksCnt < 1000 {
			fmt.Printf("; Золотых кирпичей: %d/1000", currentData.BricksCnt)
		}

		if currentData.WoodCnt < 1000 {
			fmt.Printf("; Дерева: %d/1000", currentData.WoodCnt)
		}

		fmt.Print("\n")

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
