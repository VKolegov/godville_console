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

		// this value is fixed, apparently
		// it's taken from https://godville.net/superhero
		// html element "#axe" which contains base64 text
		// decoded it looks like this
		// {"zpg":false,"d":"godville.net","p":"","td":"","u1":"wss://s2.godville.net:443/wshero","u2":"/fbh/feed?a=GjZLI9oQGPkBZMqMMMP3KYBRVcqmu"}
		// "u1" is web socket url
		// so, "u2" field is basically poll url
		pollUrl string = "https://godville.net/fbh/feed?a=GjZLI9oQGPkBZMqMMMP3KYBRVcqmu"
		pollQueryUrl string
		cnt uint = 0

		err error
	)

	for {

		pollQueryUrl = fmt.Sprintf(
			"%s&cnt=%d", pollUrl, cnt,
		)

		r, _ = eClient.Get(pollQueryUrl)

		// cleaning up
		data = structs.ExtendedData{}

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

		if lastDiaryEntry != eCurrentData.Hero.DiaryLast {
			lastDiaryEntry = eCurrentData.Hero.DiaryLast
			fmt.Printf("[Дневник] %s\n", lastDiaryEntry)
		}

		if lastNewsFromTheField != eCurrentData.NewsFromField.Msg {
			fmt.Println(eCurrentData.NewsFromField.Msg)
			lastNewsFromTheField = eCurrentData.NewsFromField.Msg
		}

		processAutoInfluence()

		if eCurrentData.Hero.TempleCompletedAt == "" {
			trackBricksExtended()
		}

		if eCurrentData.Hero.ArkCompletedAt == "" {
			trackWoodExtended()
		}

		if eCurrentData.Hero.Retirement != "" {
			trackSavingsExtended()
		}

		cnt++
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
		fmt.Printf(
			"[Храм] %s получил %d кирпичей! Итого: %d\n",
			currentData.Name,
			diff,
			bCnt,
		)
	}

	if currentData.BricksCnt == 1000 {
		fmt.Printf("[Храм] %s собрал все кирпичи для постройки храма!\n", currentData.Name)
	}
}

func trackWoodExtended() {

	wCnt := int32(eCurrentData.Hero.WoodCnt)

	if lastWoodCnt == -1 {
		lastWoodCnt = wCnt
	} else if lastWoodCnt != wCnt {
		diff := wCnt - lastWoodCnt
		lastWoodCnt = wCnt
		fmt.Printf(
			"[Ковчег] %s получил %d досок для ковчега! Итого: %d\n",
			eCurrentData.Hero.Name,
			diff,
			eCurrentData.Hero.WoodCnt,
		)
	}

	if eCurrentData.Hero.WoodCnt == 1000 {
		fmt.Printf("[Ковчег] %s собрал дерево для постройки ковчега!\n", eCurrentData.Hero.Name)
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

		fmt.Printf(
			"[Сбережения] %s отложил %d тысяч! Итого: %d тыс.",
			eCurrentData.Hero.Name,
			diff,
			savings,
		)
	}
}
