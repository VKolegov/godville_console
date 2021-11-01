package displaying

import (
	"fmt"
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

	sb.WriteByte(']')

	// fight

	// Идёт сражение
	if h.GetMonster() != "" {

		sb.WriteString(" Сражение с: ")
		sb.WriteString(h.GetMonster())
		sb.WriteString(" (")
		sb.WriteString(strconv.Itoa(h.GetMonsterProgress()))
		sb.WriteString("/100)")

		if p != nil && p.GetMonster() == h.GetMonster() {
			appendDiff(h.GetMonsterProgress(), p.GetMonsterProgress(), &sb)
		}
	}

	// health
	health := h.GetHealth()
	sb.WriteString(" Здоровье: ")
	sb.WriteString(strconv.Itoa(health))
	sb.WriteByte('/')
	sb.WriteString(strconv.Itoa(h.GetMaxHealth()))

	if p != nil {
		appendDiff(health, p.GetHealth(), &sb)
	}

	// gold
	sb.WriteString(" Золота: ")
	sb.WriteString(h.GetGoldApprox())

	if h.GetGold() >= 0 && p != nil {
		appendDiff(h.GetGold(), p.GetGold(), &sb)
	}
	sb.WriteByte(';')

	// inventory
	invNum := h.GetInvNum()
	sb.WriteString(" Инвентарь: ")
	sb.WriteString(strconv.Itoa(invNum))
	sb.WriteByte('/')
	sb.WriteString(strconv.Itoa(h.GetMaxInvNum()))

	if p != nil {
		appendDiff(h.GetInvNum(), p.GetInvNum(), &sb)
	}
	sb.WriteByte(';')

	sb.WriteByte('\n')

	fmt.Print(sb.String())
}
