package displaying

import (
	"fmt"
	"godville/structs"
	"strconv"
	"strings"
	"time"
)

func PrintGodStatus(g structs.Hero, clanInfo bool, p structs.Hero, datetimeLayout string) {


	sb := strings.Builder{}
	sb.Grow(256)

	if datetimeLayout != "" {
		t := time.Now()
		sb.WriteByte('[')
		sb.WriteString(t.Format(datetimeLayout))
		sb.WriteByte(']')
	}

	sb.WriteByte('[')
	sb.WriteString(g.GetGodName())
	sb.WriteString("] ")

	sb.WriteString("Прана: ")
	sb.WriteString(strconv.Itoa(g.GetGodPower()))
	sb.WriteByte('%')

	if p != nil {
		appendDiff(g.GetGodPower(), p.GetGodPower(), &sb)
	}

	if g.GetGodPowerCharges() >= 0 {
		sb.WriteString("; Зарядов праны: ")
		sb.WriteString(strconv.Itoa(g.GetGodPowerCharges()))

		if p != nil {
			appendDiff(g.GetGodPowerCharges(), p.GetGodPowerCharges(), &sb)
		}
	}

	if g.GetBricks() < 1000 {
		sb.WriteString("; Золотых кирпичей: ")
		sb.WriteString(strconv.Itoa(g.GetBricks()))

		if p != nil {
			appendDiff(g.GetBricks(), p.GetBricks(), &sb)
		}
	}

	if g.GetWood() < 1000 {
		sb.WriteString("; Дерева для ковчега: ")
		sb.WriteString(strconv.Itoa(g.GetWood()))

		if p != nil {
			appendDiff(g.GetWood(), p.GetWood(), &sb)
		}
	}

	if g.GetSavingsNum() >= 0 {
		sb.WriteString("; Сбережений: ")
		sb.WriteString(g.GetSavings())

		if p != nil {
			appendDiff(g.GetSavingsNum(), p.GetSavingsNum(), &sb)
		}
	}

	sb.WriteByte('\n')

	if clanInfo {

		sb.WriteByte('[')
		sb.WriteString(g.GetGodName())
		sb.WriteByte(']')

		sb.WriteString(" Клан: ")
		sb.WriteString(g.GetClan())
		sb.WriteString("; Должность: ")
		sb.WriteString(g.GetClanPosition())
		sb.WriteByte('\n')
	}

	fmt.Print(sb.String())
}