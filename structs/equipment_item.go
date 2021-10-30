package structs

type EquipmentItem struct {
	Name  string `json:"name"`
	Level string `json:"level"`
	Capt  string `json:"capt"` // Место (голова, ноги, талисман, etc.)
	B     uint8 `json:"b"`    // Жирный предмет
}
