package commands

import (
	"fmt"
	"godville/structs"
	"strings"
)

func Hero(data structs.GodvilleData) {
	fmt.Printf("Герой: %s; Уровень: %d; Характер: %s", data.Name, data.Level, data.Alignment)

	if data.Aura != "" {
		fmt.Printf("; Аура: %s", data.Aura)
	}

	fmt.Print("\n")
}

func Inventory(data structs.GodvilleData) {
	fmt.Printf("Инвентарь: %d/%d;", data.InventoryNum, data.InventoryMaxNum)

	if len(data.Activatables) == 0 {
		fmt.Println()
		return
	}

	a := strings.Join(data.Activatables, ", ")
	fmt.Printf("Активируемые вещи: %s\n", a)
}
