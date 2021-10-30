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
