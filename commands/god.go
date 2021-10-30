package commands

import (
	"fmt"
	"godville/structs"
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

func GodInfoExtended(eData *structs.ExtendedData, clanInfo bool)  {

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
