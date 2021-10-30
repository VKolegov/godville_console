package structs

type Equipment struct {
	Weapon   EquipmentItem `json:"weapon"`
	Shield   EquipmentItem `json:"shield"`
	Head     EquipmentItem `json:"head"`
	Body     EquipmentItem `json:"body"`
	Arms     EquipmentItem `json:"arms"`
	Legs     EquipmentItem `json:"legs"`
	Talisman EquipmentItem `json:"talisman"`
}
