package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"godville/commands"
	"godville/structs"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

func login(username, password string) error {
	loginData := url.Values{
		"username": {username},
		"password": {password},
	}

	cookieJar, _ := cookiejar.New(nil)

	eClient = &http.Client{
		Jar: cookieJar,
	}

	_, err := eClient.PostForm("https://godville.net/login/login", loginData)

	if err != nil {
		return errors.New(
			fmt.Sprintf("Login error: %s\n", err.Error()),
		)
	}

	return nil
}

func trackExtended(rate int) {

	var (
		r *http.Response

		initialRequest = true
		data           structs.ExtendedData

		err error
	)

	for {

		d := url.Values{
			"a": {"GjZLI9oQGPkBZMqMMMP3KYBRVcqmu"},
		}

		r, _ = eClient.PostForm("https://godville.net/fbh/feed", d)

		err = json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			fmt.Printf("Error while decoding extended data: %s", err.Error())
		}

		eCurrentData = &data

		if initialRequest {
			greetingsExtended()
			initialRequest = false
		}

		trackGodDataExtended()
		trackHeroDataExtended()
		trackFight()

		if eCurrentData.Hero.TempleCompletedAt == "" {
			trackBricksExtended()
		}

		if eCurrentData.Hero.ArkCompletedAt == "" {
			trackWoodExtended()
		}

		if eCurrentData.Hero.Retirement != "" {
			trackSavingsExtended()
		}

		time.Sleep(time.Second * time.Duration(rate))
	}
}

func greetingsExtended() {

	var (
		godvilleTimeLayout = "2006-01-02T15:04:05-07:00" // почти ISO8601. ISO8601:"-07:00" godville:"-07:00"
		dateFormat         = "2006/01/02"
		timezone           = "Asia/Krasnoyarsk" // TODO: detect from system
		loc                *time.Location

		err error
	)

	fmt.Printf("%s на связи!\n", eCurrentData.Hero.Godname)

	loc, err = time.LoadLocation(timezone)

	if err != nil {
		fmt.Printf("err: %s\n", err.Error())
		return
	}

	if eCurrentData.Hero.TempleCompletedAt != "" {

		templeDate, _ := time.ParseInLocation(
			godvilleTimeLayout,
			eCurrentData.Hero.TempleCompletedAt,
			loc,
		)

		fmt.Printf(
			"Храм достроен %s, поздравляем!\n",
			templeDate.Format(dateFormat),
		)
	}

	if eCurrentData.Hero.ArkCompletedAt != "" {
		templeDate, _ := time.ParseInLocation(
			godvilleTimeLayout,
			eCurrentData.Hero.ArkCompletedAt,
			loc,
		)

		fmt.Printf(
			"Ковчег достроен %s, поздравляем!\n",
			templeDate.Format(dateFormat),
		)
	}

	commandList()
}

func trackGodDataExtended() {

	if lastPrana != eCurrentData.Hero.Godpower ||
		(lastBrickCnt >= 0 && lastBrickCnt != int16(eCurrentData.Hero.BricksCnt)) ||
		(lastWoodCnt >= 0 && lastWoodCnt != int32(eCurrentData.Hero.WoodCnt)) ||
		lastSavingsString != eCurrentData.Hero.Retirement {

		commands.GodInfoExtended(eCurrentData, false)

		lastPrana = eCurrentData.Hero.Godpower
		lastSavingsString = eCurrentData.Hero.Retirement
	}
}

func trackHeroDataExtended() {
	var (
		whereabouts string
	)

	if lastDiaryEntry != eCurrentData.Hero.DiaryLast {
		lastDiaryEntry = eCurrentData.Hero.DiaryLast
		fmt.Printf("[Дневник] %s\n", lastDiaryEntry)
	}

	fmt.Println(eCurrentData.NewsFromField.Msg)

	if lastHealth != eCurrentData.Hero.Health ||
		lastPillar != eCurrentData.Hero.Distance ||
		lastTown != eCurrentData.Hero.TownName ||
		lastGold != eCurrentData.Hero.GoldWe {

		if eCurrentData.Hero.TownName == "" {
			whereabouts = fmt.Sprintf("Столб #%d", eCurrentData.Hero.Distance)
		} else {
			whereabouts = fmt.Sprintf("%s (ст. %d)", eCurrentData.Hero.TownName, eCurrentData.Hero.Distance)
		}

		fmt.Printf(
			"[%s] Здоровье: %d/%d; Золота: %s; Инвентарь: %d/%d\n",
			whereabouts,
			eCurrentData.Hero.Health,
			eCurrentData.Hero.MaxHealth,
			eCurrentData.Hero.GoldWe,
			eCurrentData.Hero.InventoryNum,
			eCurrentData.Hero.InventoryMaxNum,
		)

		lastHealth = eCurrentData.Hero.Health
		lastTown = eCurrentData.Hero.TownName
		lastPillar = eCurrentData.Hero.Distance
		lastGold = eCurrentData.Hero.GoldWe
	}
}

func trackFight() {
	// Идёт сражение
	if eCurrentData.Hero.MonsterName != "" {

		if lastMonsterProgress != eCurrentData.Hero.MonsterProgress {
			fmt.Printf(
				"%s сражается с %s (%d/100)\n",
				eCurrentData.Hero.Name,
				eCurrentData.Hero.MonsterName,
				eCurrentData.Hero.MonsterProgress,
			)

			lastMonsterProgress = eCurrentData.Hero.MonsterProgress
		}

		fmt.Println(eCurrentData.NewsFromField.Msg)

	} else {
		lastMonsterProgress = 0
	}
}

func trackBricksExtended() {

	bCnt := int16(currentData.BricksCnt)

	if lastWoodCnt == -1 {
		lastBrickCnt = bCnt
	} else if lastBrickCnt != bCnt {
		diff := bCnt - lastBrickCnt
		lastBrickCnt = bCnt
		fmt.Printf("%s получил %d кирпичей! Итого: %d\n", currentData.Name, diff, bCnt)
	}

	if currentData.BricksCnt == 1000 {
		fmt.Printf("%s собрал все кирпичи для постройки храма!\n", currentData.Name)
	}
}

func trackWoodExtended() {

	wCnt := int32(eCurrentData.Hero.WoodCnt)

	if lastWoodCnt == -1 {
		lastWoodCnt = wCnt
	} else if lastWoodCnt != wCnt {
		diff := wCnt - lastWoodCnt
		lastWoodCnt = wCnt
		fmt.Printf("Герой получил %d досок для ковчега!\n", diff)
	}

	if eCurrentData.Hero.WoodCnt == 1000 {
		fmt.Printf("%s собрал дерево для постройки ковчега!\n", eCurrentData.Hero.Name)
	}
}

func trackSavingsExtended() {

	savings, err := parseSavings(eCurrentData.Hero.Retirement)

	if err != nil {
		fmt.Println(err)
		return
	}

	if lastSavings == -1 {
		lastSavings = savings
	} else if lastSavings != savings {

		diff := savings - lastSavings
		lastSavings = savings

		fmt.Printf("Герой отложил %d тысяч! Итого: %d тыс.", diff, savings)
	}
}
