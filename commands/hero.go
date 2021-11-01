package commands

import (
	"fmt"
	"github.com/fatih/color"
	"godville/displaying"
	"godville/structs"
	"net/http"
	"strconv"
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
		eData.Hero.ArenaLost,
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

	sb := strings.Builder{}

	sb.WriteString("Инвентарь: ")

	sb.WriteString(displaying.InventoryColor.Sprint(" "))
	sb.WriteString(displaying.InventoryColor.Sprint(strconv.Itoa(int(eData.Hero.InventoryNum))))
	sb.WriteString(displaying.InventoryColor.Sprint("/"))
	sb.WriteString(displaying.InventoryColor.Sprint(strconv.Itoa(int(eData.Hero.InventoryMaxNum))))
	sb.WriteString(displaying.InventoryColor.Sprint(" "))
	sb.WriteByte('\n')

	fmt.Print(sb.String())

	if len(eData.Inventory) == 0 {
		return
	}

	fmt.Println("На себе герой несёт:")
	for itemName, item := range eData.Inventory {

		fmt.Printf("   - [%d] ", item.Pos)

		if itemName == "золотой кирпич" {
			displaying.GoldColor.Set()
		}

		if item.Price > 100 {
			color.Set(color.Bold)
		}

		if item.Type == "heal_potion" {
			color.Set(color.FgGreen)
		}

		if item.ActivateByUser {
			fmt.Print("@ ")
		}

		fmt.Printf("%s", itemName)
		if item.Cnt > 1 {
			fmt.Printf(" (%d шт.)", item.Cnt)
		}

		if item.Price >= 100 {
			fmt.Print(" (ценн.)")
		}
		if item.Type == "heal_potion" {
			fmt.Print(" (леч.)")
		}

		if item.Description != "" {
			fmt.Printf("  -  %s", item.Description)
		}

		if item.ActivateByUser {
			fmt.Printf(" (цена: %d праны)", item.NeedsGodpower)
		}

		color.Set(color.Reset)
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

		found bool

		err error
	)

	for itemName, item = range d.Inventory {
		if item.Pos == id {
			found = true
			break
		}
	}

	if !found {
		displaying.InventoryColor.Printf("[Инвентарь] Предмет с таким ID не найден")
		fmt.Print("\n")
		return
	}

	if item.ActivateByUser == false {
		displaying.InventoryColor.Printf("[Инвентарь] %s нельзя активировать", itemName)
		fmt.Print("\n")
		return
	}

	if d.Hero.Godpower < item.NeedsGodpower {
		displaying.InventoryColor.Printf(
			"[Инвентарь] Не хватает силёнок чтобы активировать %s (%d/%d)",
			itemName,
			d.Hero.Godpower,
			item.NeedsGodpower,
		)
		fmt.Print("\n")
		return
	}

	rData := map[string]interface{}{
		"id": itemName,
	}

	_, err = MakeFeedPostRequest(c, "agQHqM4rCoT0CaDvq44I", rData)

	if err != nil {
		displaying.InventoryColor.Printf(
			"[Инвентарь] Не смогли активировать %s, причина: %s",
			itemName,
			err.Error(),
		)
		fmt.Print("\n")
		return
	}

	displaying.InventoryColor.Printf("[Инвентарь] Активировали %s!", itemName)
	fmt.Print("\n")
}
