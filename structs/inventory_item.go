package structs

type InventoryItem struct {
	Cnt            int    `json:"cnt"`
	Pos            int    `json:"pos"`
	Price          int    `json:"price"`
	Type           string `json:"type"`
	ActivateByUser bool   `json:"activate_by_user"`
	NeedsGodpower  uint8    `json:"needs_godpower"`
	Description    string `json:"description"`
}
