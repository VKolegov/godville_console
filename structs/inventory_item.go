package structs

type InventoryItem struct {
	Cnt   int    `json:"cnt"`
	Pos   int    `json:"pos"`
	Price int    `json:"price"`
	Type  string `json:"type"`
}
