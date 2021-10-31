package commands

import (
	"fmt"
	"godville/structs"
	"net/http"
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
		return
	}

	fmt.Println("На себе герой несёт:")
	for itemName, item := range eData.Inventory {

		fmt.Print("   - ")

		if item.ActivateByUser {
			fmt.Print("@ ")
		}

		fmt.Printf("[%d] %s (%d шт.)", item.Pos, itemName, item.Cnt)

		if item.Type == "heal_potion" {
			fmt.Print(" (лечебн.)")
		}

		if item.Description != "" {
			fmt.Printf("  -  %s", item.Description)
		}

		if item.ActivateByUser {
			fmt.Printf(" (цена: %d праны)", item.NeedsGodpower)
		}

		fmt.Print("\n")
	}
}

func Equipment(eData *structs.ExtendedData) {
	fmt.Println("Снаряжение героя:")

	printEquipmentItem(eData.Equipment.Head)
	printEquipmentItem(eData.Equipment.Talisman)
	printEquipmentItem(eData.Equipment.Body)
	printEquipmentItem(eData.Equipment.Arms)
	printEquipmentItem(eData.Equipment.Weapon)
	printEquipmentItem(eData.Equipment.Shield)
	printEquipmentItem(eData.Equipment.Legs)
}

func printEquipmentItem(item structs.EquipmentItem) {
	fmt.Printf("	- ")
	if item.B == 1 {
		fmt.Print("[Ж]")
	}
	fmt.Printf("[%s] %s %s\n", item.Capt, item.Name, item.Level)
}

func UseItem(id int, d *structs.ExtendedData, c *http.Client) {
	var (
		itemName string
		item     structs.InventoryItem

		err error
	)

	for itemName, item = range d.Inventory {
		if item.Pos == id {
			break
		}
	}

	if item.ActivateByUser == false {
		fmt.Printf("[Инвентарь] %s нельзя активировать\n", itemName)
		return
	}

	if d.Hero.Godpower < item.NeedsGodpower {
		fmt.Printf(
			"[Инвентарь] Не хватает силёнок чтобы активировать %s (%d/%d)\n",
			itemName,
			d.Hero.Godpower,
			item.NeedsGodpower,
		)
		return
	}

	rData := map[string]interface{}{
		"id": itemName,
	}

	_, err = MakeFeedPostRequest(c, "agQHqM4rCoT0CaDvq44I", rData)

	if err != nil {
		fmt.Printf("[Инвентарь] Не смогли активировать %s, причина: %s", itemName, err.Error())
		return
	}

	fmt.Printf("[Инвентарь] Активировали %s!\n", itemName)
}
