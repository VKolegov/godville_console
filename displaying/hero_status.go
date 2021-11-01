package displaying

import (
	"fmt"
	"github.com/fatih/color"
	"godville/structs"
	"strconv"
	"strings"
	"time"
)

func PrintHeroStatus(h structs.Hero, p structs.Hero, datetimeLayout string) {

	var sb strings.Builder

	sb.Grow(256)

	if datetimeLayout != "" {
		t := time.Now()
		sb.WriteByte('[')
		sb.WriteString(t.Format(datetimeLayout))
		sb.WriteByte(']')
	}

	sb.WriteByte('[')

	pillarStr := strconv.Itoa(h.GetPillar())

	town := h.GetTown()
	if town == "" {
		sb.WriteString("Столб #")
		sb.WriteString(pillarStr)
	} else {
		sb.WriteString(town)
		sb.WriteString(" (ст. ")
		sb.WriteString(pillarStr)
		sb.WriteByte(')')
	}

	sb.WriteString("] ")

	// health
	health := h.GetHealth()
	sb.WriteString("Здоровье: ")

	if health < 25 {
		HealthColor.Add(color.BlinkSlow)
	} else {
		resetColors()
	}

	sb.WriteString(HealthColor.Sprint(" "))
	sb.WriteString(HealthColor.Sprint(strconv.Itoa(health)))
	sb.WriteString(HealthColor.Sprint("/"))
	sb.WriteString(HealthColor.Sprint(strconv.Itoa(h.GetMaxHealth())))
	sb.WriteString(HealthColor.Sprint(" "))

	if p != nil {
		appendDiff(health, p.GetHealth(), &sb)
	}

	sb.WriteString(" | ")

	// fight
	if h.GetMonster() != "" {

		sb.WriteString("Противник: ")

		var monsterColor *color.Color

		if h.IsMonsterTough() {
			monsterColor = ToughMonsterColor
		} else {
			monsterColor = RegularMonsterColor
		}

		sb.WriteString(monsterColor.Sprint(h.GetMonster()))
		sb.WriteString(monsterColor.Sprint(" ("))
		sb.WriteString(monsterColor.Sprint(strconv.Itoa(h.GetMonsterProgress())))
		sb.WriteString(monsterColor.Sprint("/100)"))

		if p != nil && p.GetMonster() == h.GetMonster() {
			appendDiff(h.GetMonsterProgress(), p.GetMonsterProgress(), &sb)
		}

		sb.WriteString(" | ")
	}

	// gold
	sb.WriteString("Золота: ")

	sb.WriteString(GoldColor.Sprint(" "))
	sb.WriteString(GoldColor.Sprint(h.GetGoldApprox()))
	sb.WriteString(GoldColor.Sprint(" "))

	if h.GetGold() >= 0 && p != nil {
		appendDiff(h.GetGold(), p.GetGold(), &sb)
	}
	sb.WriteString(" | ")

	// inventory
	invNum := h.GetInvNum()
	sb.WriteString("Инвентарь: ")

	sb.WriteString(InventoryColor.Sprint(" "))
	sb.WriteString(InventoryColor.Sprint(strconv.Itoa(invNum)))
	sb.WriteString(InventoryColor.Sprint("/"))
	sb.WriteString(InventoryColor.Sprint(strconv.Itoa(h.GetMaxInvNum())))
	sb.WriteString(InventoryColor.Sprint(" "))

	if p != nil {
		appendDiff(h.GetInvNum(), p.GetInvNum(), &sb)
	}

	sb.WriteByte('\n')

	fmt.Print(sb.String())
}
