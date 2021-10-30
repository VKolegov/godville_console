package commands

import (
	"fmt"
	"godville/structs"
)

func QuestStatus(data structs.GodvilleData) {
	fmt.Printf("Текущий квест: \"%s\"; прогресс: %d%%\n", data.Quest, data.QuestProgress)
}

func QuestStatusExtended(hero structs.Hero) {
	fmt.Printf(
		"Текущий квест: \"%s\"; прогресс: %d%%; Выполнено квестов: %d\n",
		hero.Quest,
		hero.QuestProgress,
		hero.QuestsCompleted,
	)
}
