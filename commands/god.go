package commands

import (
	"encoding/json"
	"fmt"
	"godville/enc"
	"godville/structs"
	"net/http"
	"net/url"
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

func MakeEvil(eClient *http.Client) {

	var (
		r *http.Response

		inf structs.Influence

		err error
	)

	rData := map[string]interface{}{
		"action": "punish",
		//"confirm": "1", // could be present, maybe it has something to do with arena
		//"cid":     nil, // could be present, maybe it has something to do with arena
		//"s":       nil, // could be present, maybe it has something to do with arena
	}

	rDataEncoded, err := json.Marshal(rData)

	if err != nil {
		fmt.Printf("Error while encoding punish request: %s\n", err)
	}

	a := enc.Vm("5JgMUahE1BYdtf7quoWz")
	b := enc.Wm(rDataEncoded)

	d := url.Values{
		"a": {a}, // e.g. kJFiYFQT8EtYAQwiIgmiUA2VWngYQ
		"b": {b}, // e.g. W0vFCeyJhY3Rpb24iOiJwdW5pc2gifQ==GrS
	}

	r, _ = eClient.PostForm("https://godville.net/fbh/feed", d)

	err = json.NewDecoder(r.Body).Decode(&inf)

	if err != nil {
		fmt.Printf("Ошибка при попытке совершить зло: %s", err.Error())
		return
	}

	if inf.Status == "success" {
		fmt.Println("[попытка зла засчитана]")
	} else {
		fmt.Println("[попытка зла не засчитана]")
		fmt.Printf("%+v\n", inf)
	}

	fmt.Printf("[влияние:зло] %s\n", inf.DisplayString)
}
