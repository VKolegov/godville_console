package structs

import "godville/utils"

type GodvilleData struct {
	Name              string   `json:"name"`
	Godname           string   `json:"godname"`
	Gender            string   `json:"gender"`
	Level             uint16   `json:"level"`
	MaxHealth         uint16   `json:"max_health"`
	InventoryMaxNum   uint16   `json:"inventory_max_num"`
	Motto             string   `json:"motto"`
	Clan              string   `json:"clan"`
	ClanPosition      string   `json:"clan_position"`
	Alignment         string   `json:"alignment"`
	BricksCnt         int      `json:"bricks_cnt"`
	WoodCnt           uint32   `json:"wood_cnt"`
	TempleCompletedAt string   `json:"temple_completed_at"`
	Pet               Pet      `json:"pet"`
	ArkCompletedAt    string   `json:"ark_completed_at"`
	ArenaWon          uint     `json:"arena_won"`
	ArenaLost         uint     `json:"arena_lost"`
	Savings           string   `json:"savings"`
	Health            uint16   `json:"health"`
	QuestProgress     uint16   `json:"quest_progress"`
	ExpProgress       uint16   `json:"exp_progress"`
	Expired           bool     `json:"expired"`
	Godpower          uint8    `json:"godpower"`
	GoldApprox        string   `json:"gold_approx"`
	DiaryLast         string   `json:"diary_last"`
	TownName          string   `json:"town_name"`
	Distance          uint16   `json:"distance"`
	ArenaFight        bool     `json:"arena_fight"`
	InventoryNum      uint16   `json:"inventory_num"`
	Quest             string   `json:"quest"`
	Aura              string   `json:"aura"`
	Activatables      []string `json:"activatables"`
}

func (g GodvilleData) GetName() string {
	return g.Name
}

func (g GodvilleData) GetHealth() int {
	return int(g.Health)
}

func (g GodvilleData) GetMaxHealth() int {
	return int(g.MaxHealth)
}

func (g GodvilleData) GetInvNum() int {
	return int(g.InventoryNum)
}

func (g GodvilleData) GetMaxInvNum() int {
	return int(g.InventoryMaxNum)
}

func (g GodvilleData) GetPillar() int {
	return int(g.Distance)
}

func (g GodvilleData) GetTown() string {
	return g.TownName
}

func (g GodvilleData) GetGold() int {
	return -1
}

func (g GodvilleData) GetGoldApprox() string {
	return g.GoldApprox
}

func (g GodvilleData) GetGodName() string {
	return g.Godname
}

func (g GodvilleData) GetGodPower() int {
	return int(g.Godpower)
}

func (g GodvilleData) GetGodPowerCharges() int {
	return -1
}

func (g GodvilleData) GetSavings() string {
	return g.Savings
}

func (g GodvilleData) GetSavingsNum() int {

	if g.GetSavings() == "" {
		return 0
	}

	s, _ := utils.ParseSavings(g.GetSavings())
	return s
}

func (g GodvilleData) GetBricks() int {
	return g.BricksCnt
}

func (g GodvilleData) GetWood() int {
	return int(g.WoodCnt)
}

func (g GodvilleData) GetClan() string {
	return g.Clan
}

func (g GodvilleData) GetClanPosition() string {
	return g.ClanPosition
}
