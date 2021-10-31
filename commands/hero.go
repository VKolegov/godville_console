package commands

import (
	"encoding/json"
	"fmt"
	"godville/enc"
	"godville/structs"
	"net/http"
	"net/url"
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

func UseItem(id int, inventory map[string]structs.InventoryItem, c *http.Client) {
	var (
		r *http.Response

		response structs.GenericResponse
		itemName string
		item     structs.InventoryItem

		err error
	)

	for itemName, item = range inventory {
		if item.Pos == id {
			break
		}
	}

	rData := map[string]string{
		"id": itemName,
	}

	rDataEncoded, err := json.Marshal(rData)

	if err != nil {
		fmt.Printf("Error while encoding item request: %s\n", err)
	}

	a := enc.Vm("agQHqM4rCoT0CaDvq44I")
	b := enc.Wm(rDataEncoded)

	d := url.Values{
		"a": {a}, // e.g. 9FwH2ahcM6oMrfS4DfuMyv1gcJksp
		"b": {b}, // e.g. DvApzeyJpZCI6ItGB0LLQtdGC0Y/RidGD0Y7RgdGPINGC0YvQutCy0YMifQ==9is // светящуюся тыкву
	}

	fmt.Printf("req: %+v,\n", d)

	r, _ = c.PostForm("https://godville.net/fbh/feed", d)

	err = json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		fmt.Printf("Ошибка при попытке распознать результат исп предмета: %s", err.Error())
		return
	}

	if response.Status != "success" {
		fmt.Println("[не удалось донести запрос до сервера]")
		fmt.Printf("%+v\n", response)
	}

	fmt.Printf("[Инвентарь] Предмет:%s активирован...\n", itemName)
}
