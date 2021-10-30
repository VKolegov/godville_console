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

	kd := float64(data.ArenaWon) / float64(data.ArenaLost)
	fmt.Printf("Побед на арене: %d; Поражений на арене: %d; К/Д: %.2f\n", data.ArenaWon, data.ArenaLost, kd)
}

func HeroExtended(eData *structs.ExtendedData) {
	fmt.Printf(
		"Герой: %s; Уровень: %d; Характер: %s",
		eData.Hero.Name,
		eData.Hero.Level,
		eData.Hero.Alignment,
	)

	if eData.Hero.AuraName != "" {
		fmt.Printf("; Аура: %s (%d сек)", eData.Hero.AuraName, eData.Hero.AuraTime)
	}

	fmt.Print("\n")

	kd := float64(eData.Hero.MonstersKilled) / float64(eData.Hero.DeathCount)

	fmt.Printf(
		"Убито монстров: %d; Смертей: %d; K/D: %.2f\n",
		eData.Hero.MonstersKilled,
		eData.Hero.DeathCount,
		kd,
	)

	kd = float64(eData.Hero.ArenaWon) / float64(eData.Hero.ArenaLost)
	fmt.Printf(
		"Побед на арене: %d; Поражений на арене: %d; K/D: %.2f\n",
		eData.Hero.ArenaWon,
		eData.Hero.ArenaWon,
		kd,
	)
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

func InventoryExtended(eData *structs.ExtendedData) {
	fmt.Printf("Инвентарь: %d/%d\n", eData.Hero.InventoryNum, eData.Hero.InventoryMaxNum)

	if len(eData.Inventory) == 0 {
		fmt.Println()
		return
	}

	fmt.Println("На себе герой несёт:")
	for itemName, item := range eData.Inventory {
		fmt.Printf("%s (%d шт.)", itemName, item.Cnt)

		if item.Type == "heal_potion" {
			fmt.Printf(" (лечебн.)")
		}

		fmt.Print("\n")
	}
}
