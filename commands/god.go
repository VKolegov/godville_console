package commands

import (
	"encoding/json"
	"fmt"
	"godville/structs"
	"net/http"
	"strconv"
	"strings"
)

func PrintGodInfo(g structs.Hero, clanInfo bool) {

	sb := strings.Builder{}
	sb.Grow(120)

	sb.WriteByte('[')
	sb.WriteString(g.GetGodName())
	sb.WriteString("] ")

	sb.WriteString("Прана: ")
	sb.WriteString(strconv.Itoa(g.GetGodPower()))
	sb.WriteByte('%')

	if g.GetGodPowerCharges() >= 0 {
		sb.WriteString(" (зарядов: ")
		sb.WriteString(strconv.Itoa(g.GetGodPowerCharges()))
		sb.WriteByte(')')
	}

	if g.GetBricks() < 1000 {
		sb.WriteString("; Золотых кирпичей: ")
		sb.WriteString(strconv.Itoa(g.GetBricks()))
	}

	if g.GetWood() < 1000 {
		sb.WriteString("; Дерева для ковчега: ")
		sb.WriteString(strconv.Itoa(g.GetWood()))
	}

	if g.GetSavings() != "" {
		sb.WriteString("; Сбережений: ")
		sb.WriteString(g.GetSavings())
	}

	sb.WriteByte('\n')

	if clanInfo {

		sb.WriteByte('[')
		sb.WriteString(g.GetGodName())
		sb.WriteByte(']')

		sb.WriteString(" Клан: ")
		sb.WriteString(g.GetClan())
		sb.WriteString("; Должность: ")
		sb.WriteString(g.GetClanPosition())
		sb.WriteByte('\n')
	}

	fmt.Print(sb.String())
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
