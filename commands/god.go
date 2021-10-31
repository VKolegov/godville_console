package commands

import (
	"encoding/json"
	"fmt"
	"godville/structs"
	"net/http"
)

func GodInfo(data structs.GodvilleData) {

	fmt.Printf("[%s] Прана: %d%%", data.Godname, data.Godpower)

	if data.Savings != "" {
		fmt.Printf("; Сбережений: %s", data.Savings)
	}

	if data.BricksCnt < 1000 {
		fmt.Printf("; Золотых кирпичей: %d/1000", data.BricksCnt)
	}

	if data.WoodCnt < 1000 {
		fmt.Printf("; Дерева: %d/1000", data.WoodCnt)
	}

	fmt.Print("\n")
}

func GodInfoExtended(eData *structs.ExtendedData, clanInfo bool) {

	fmt.Printf(
		"[%s] Прана: %d%% (зарядов: %.0f)",
		eData.Hero.Godname,
		eData.Hero.Godpower,
		eData.Hero.Accumulator,
	)

	if eData.Hero.Retirement != "" {
		fmt.Printf("; Сбережений: %s", eData.Hero.Retirement)
	}

	if eData.Hero.BricksCnt < 1000 {
		fmt.Printf("; Золотых кирпичей: %d/1000", eData.Hero.BricksCnt)
	}

	if eData.Hero.WoodCnt < 1000 {
		fmt.Printf("; Дерева: %d/1000", eData.Hero.WoodCnt)
	}

	fmt.Print("\n")

	if clanInfo {
		fmt.Printf(
			"[%s] Клан: %s; Должность: %s\n",
			eData.Hero.Godname,
			eData.Hero.Clan,
			eData.Hero.ClanPosition,
		)
	}
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
		fmt.Printf("На %s, увы, силёнок не хватает\n", influenceName)
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
		fmt.Printf("[Влияние:%s] Не удалось повлиять. Причина: %s", influenceName, err.Error())
		return
	}

	err = json.Unmarshal(responseBody, &influenceResponse)

	if err != nil {
		fmt.Printf("[Влияние:%s] Не удалось распознать результат влияния: %s", influenceName, err.Error())
		return
	}

	fmt.Printf("[Влияние:%s] %s\n", influenceName, influenceResponse.DisplayString)
}

func ResurrectHero(c *http.Client, d *structs.ExtendedData) {
	var (
		err error
	)

	if d.Hero.Health > 0 {
		fmt.Printf("%s здоров как бык!... По крайней мере, ещё дышит\n", d.Hero.Name)
		return
	}

	rData := map[string]interface{}{
		"action": "resurrect",
	}

	_, err = MakeFeedPostRequest(c, "5JgMUahE1BYdtf7quoWz", rData)

	if err != nil {
		fmt.Printf("[Оживление] Ошибка при попытке оживить героя: %s\n", err.Error())
		return
	}

	fmt.Printf("[Оживление] Герой оживлён!\n")
}
