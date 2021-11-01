package commands

import (
	"encoding/json"
	"fmt"
	"godville/displaying"
	"godville/structs"
	"net/http"
)

func PrintGodInfo(g structs.Hero, clanInfo bool) {
	displaying.PrintGodStatus(g, clanInfo, nil, "")
}

func MakeInfluence(influenceType string, eData *structs.ExtendedData, eClient *http.Client) {
	var (
		influenceResponse structs.Influence
		influenceName     string

		responseBody []byte

		err error
	)

	switch influenceType {
	case "punish":
		influenceName = "зло"
	case "encourage":
		influenceName = "добро"
	default:
		return
	}

	if eData.Hero.Godpower < 25 {
		displaying.PranaColor.Printf("На %s, увы, силёнок не хватает", influenceName)
		fmt.Print("\n")
		return
	}

	rData := map[string]interface{}{
		"action": influenceType,
		//"confirm": "1", // could be present, maybe it has something to do with arena
		//"cid":     nil, // could be present, maybe it has something to do with arena
		//"s":       nil, // could be present, maybe it has something to do with arena
	}

	responseBody, err = MakeFeedPostRequest(eClient, "5JgMUahE1BYdtf7quoWz", rData)

	if err != nil {
		displaying.PranaColor.Printf(
			"[Влияние:%s] Не удалось повлиять. Причина: %s",
			influenceName,
			err.Error(),
		)
		return
	}

	err = json.Unmarshal(responseBody, &influenceResponse)

	if err != nil {
		displaying.PranaColor.Printf(
			"[Влияние:%s] Не удалось распознать результат влияния: %s",
			influenceName,
			err.Error(),
		)
		return
	}

	displaying.PranaColor.Printf("[Влияние:%s] %s", influenceName, influenceResponse.DisplayString)
	fmt.Print("\n")
}

func ResurrectHero(c *http.Client, d *structs.ExtendedData) {
	var (
		err error
	)

	if d.Hero.Health > 0 {
		displaying.HealthColor.Printf("%s здоров как бык!... По крайней мере, ещё дышит", d.Hero.Name)
		fmt.Print("\n")
		return
	}

	rData := map[string]interface{}{
		"action": "resurrect",
	}

	_, err = MakeFeedPostRequest(c, "5JgMUahE1BYdtf7quoWz", rData)

	if err != nil {
		displaying.HealthColor.Printf("[Оживление] Ошибка при попытке оживить героя: %s", err.Error())
		fmt.Print("\n")
		return
	}

	displaying.HealthColor.Printf("[Оживление] %s оживлён!", d.Hero.Name)
	fmt.Print("\n")
}

func GodPhrase(phrase string, c *http.Client, d *structs.ExtendedData) {
	var (
		err error
	)

	if d.Hero.Godpower < 5 {
		displaying.PranaColor.Printf("Извини, %s, ты слишком слаб и пока помолчишь...", d.Hero.Godname)
		fmt.Print("\n")
		return
	}

	rData := map[string]interface{}{
		"action": "god_phrase",
		"god_phrase": phrase,
	}

	_, err = MakeFeedPostRequest(c, "5JgMUahE1BYdtf7quoWz", rData)

	if err != nil {
		displaying.PranaColor.Printf("[Глас] Ошибка при попытке произнести глас: %s", err.Error())
		fmt.Print("\n")
		return
	}

	displaying.PranaColor.Printf("[Глас] Да услышит же %s твои сладкие речи!", d.Hero.Name)
	fmt.Print("\n")
}