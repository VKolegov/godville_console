package main

import (
	"fmt"
	"godville/commands"
	"os"
	"strconv"
	"strings"
)

func processAutoInfluence() {
	processAutoEvilInfluence()
}

func processAutoEvilInfluence() {

	var (
		autoEvilInfluenceThreshold int
		monsterProgressThreshold   int

		err error
	)

	autoEvilInfluenceThresholdStr := strings.TrimSpace(
		os.Getenv("AUTO_EVIL_INFLUENCE_THRESHOLD"),
	)

	monsterProgressThresholdStr := strings.TrimSpace(
		os.Getenv("MONSTER_PROGRESS_THRESHOLD"),
	)

	if autoEvilInfluenceThresholdStr == "" {
		return
	}

	autoEvilInfluenceThreshold, err = strconv.Atoi(autoEvilInfluenceThresholdStr)

	if err != nil {
		fmt.Printf("Ошибка при парсинге AUTO_EVIL_INFLUENCE_THRESHOLD: %s\n", err.Error())
		return
	}

	if monsterProgressThresholdStr != "" {
		monsterProgressThreshold, err = strconv.Atoi(autoEvilInfluenceThresholdStr)

		if err != nil {
			fmt.Printf("Ошибка при парсинге MONSTER_PROGRESS_THRESHOLD: %s\n", err.Error())
			fmt.Println("MONSTER_PROGRESS_THRESHOLD: используем значение по-умолчанию: 50%")
			monsterProgressThreshold = 50
		}
	} else {

		if eCurrentData.Hero.Godpower >= uint8(autoEvilInfluenceThreshold) {
			fmt.Println("[auto] Делаем автоматическое зло...")
			commands.MakeEvil(eClient)
			return
		}
	}

	// Если в бою, хватает праны и прогресс битвы не больше чем указанный
	if eCurrentData.Hero.MonsterName != "" &&
		eCurrentData.Hero.Godpower >= uint8(autoEvilInfluenceThreshold) &&
		eCurrentData.Hero.MonsterProgress <= uint16(monsterProgressThreshold) {

		fmt.Println("[auto] Делаем автоматическое зло в бою...")
		commands.MakeEvil(eClient)

	}

}

// TODO: activating item
// eg: светящуюся тыкву. позиция: 9
// a: 9FwH2ahcM6oMrfS4DfuMyv1gcJksp
// b: DvApzeyJpZCI6ItGB0LLQtdGC0Y/RidGD0Y7RgdGPINGC0YvQutCy0YMifQ==9is
// eg: оранжевую коробочку. позиция: 13
// a: JqQb2ahcM6oMrfS4DfuMyv1gLFw3k
// b: 2rikgeyJpZCI6ItC+0YDQsNC90LbQtdCy0YPRjiDQutC+0YDQvtCx0L7Rh9C60YMifQ==0XG