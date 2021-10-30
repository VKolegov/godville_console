package structs

type InventoryItem struct {
	Cnt            int    `json:"cnt"`              // Количество
	Pos            int    `json:"pos"`              // Позиция в инвентаре
	Price          int    `json:"price"`            // Цена
	Type           string `json:"type"`             // Тип предмета
	ActivateByUser bool   `json:"activate_by_user"` // Активируемый
	NeedsGodpower  uint8  `json:"needs_godpower"`   // Сколько нужно праны для активации
	Description    string `json:"description"`      // Описание
}
