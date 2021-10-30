package main

import "time"

type responseJSON struct {
	Name              string    `json:"name"`
	Godname           string    `json:"godname"`
	Gender            string    `json:"gender"`
	Level             uint16    `json:"level"`
	MaxHealth         uint16    `json:"max_health"`
	InventoryMaxNum   uint16    `json:"inventory_max_num"`
	Motto             string    `json:"motto"`
	Clan              string    `json:"clan"`
	ClanPosition      string    `json:"clan_position"`
	Alignment         string    `json:"alignment"`
	BricksCnt         int       `json:"bricks_cnt"`
	WoodCnt           int       `json:"wood_cnt"`
	TempleCompletedAt time.Time `json:"temple_completed_at"`
	Pet               struct {
		PetName  string `json:"pet_name"`
		PetClass string `json:"pet_class"`
		PetLevel string `json:"pet_level"`
	} `json:"pet"`
	ArkCompletedAt interface{} `json:"ark_completed_at"`
	ArenaWon       uint        `json:"arena_won"`
	ArenaLost      uint        `json:"arena_lost"`
	Savings        string      `json:"savings"`
	Health         uint16      `json:"health"`
	QuestProgress  uint16      `json:"quest_progress"`
	ExpProgress    uint16      `json:"exp_progress"`
	Expired        bool        `json:"expired"`
	Godpower       uint8       `json:"godpower"`
	GoldApprox     string      `json:"gold_approx"`
	DiaryLast      string      `json:"diary_last"`
	TownName       string      `json:"town_name"`
	Distance       uint16      `json:"distance"`
	ArenaFight     bool        `json:"arena_fight"`
	InventoryNum   uint16      `json:"inventory_num"`
	Quest          string      `json:"quest"`
	Activatables   []string    `json:"activatables"`
}
