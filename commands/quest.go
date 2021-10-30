package commands

import (
	"fmt"
	"godville/structs"
)

func QuestStatus(data structs.GodvilleData) {
	fmt.Printf("Текущий квест: \"%s\"; прогресс: %d%%\n", data.Quest, data.QuestProgress)
}
