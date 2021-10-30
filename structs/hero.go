package structs

type Hero struct {
	Name                string        `json:"name"`                   // Имя героя
	Godname             string        `json:"godname"`                // Имя бога
	Gender              string        `json:"gender"`                 // Пол
	GGender             string        `json:"g_gender"`               // Пол (бога?)
	GoldApprox          string        `json:"gold_approx"`            // Приблизительное кол-во золота
	Level               int           `json:"level"`                  // Уровень героя
	Quest               string        `json:"quest"`                  // Название квеста
	QuestProgress       int           `json:"quest_progress"`         // Прогресс квеста в процентах
	ExpProgress         int           `json:"exp_progress"`           // Прогресс опыта до сл. уровня (в процентах)
	Health              uint16        `json:"health"`                 // Здоровье
	MaxHealth           uint16        `json:"max_health"`             // Макс. здоровье
	InventoryNum        uint16        `json:"inventory_num"`          // Кол-во вещей в инвентаре
	InventoryMaxNum     uint16        `json:"inventory_max_num"`      // Максимальное кол-во вещей в инвентаре
	Alignment           string        `json:"alignment"`              // Характер
	Motto               string        `json:"motto"`                  // Девиз
	Clan                string        `json:"clan"`                   // Название клана
	TimeZone            string        `json:"time_zone"`              // Название города (временная зона)
	DiaryLast           string        `json:"diary_last"`             // Последняя запись в дневнике
	TempleCompletedAt   string        `json:"temple_completed_at"`    // Когда завершилось строительство храма
	MfAt                interface{}   `json:"mf_at"`                  // TODO: ?
	ArkCompletedAt      string        `json:"ark_completed_at"`       // Когда завершилось строительство ковчега
	Distance            uint16        `json:"distance"`               // Номер столба от столицы
	TownName            string        `json:"town_name"`              // Название города
	InTown              bool          `json:"in_town"`                // Находится ли герой в городе
	ArenaFight          bool          `json:"arena_fight"`            // Находится ли герой в битве (босс, арена, подземелье)
	BricksCnt           int           `json:"bricks_cnt"`             // Количество собранных золотых кирпичей
	Wood                string        `json:"wood"`                   // Количество собранной древесины для ковчега в процентах
	WoodCnt             int           `json:"wood_cnt"`               // Количество собранной древесины
	Godpower            uint8         `json:"godpower"`               // Количество праны
	ClanPosition        string        `json:"clan_position"`          // Должность в гильдии
	ACmd                bool          `json:"a_cmd"`                  // TODO: ?
	PetsMax             int           `json:"pets_max"`               // TODO: ?
	AuraName            string        `json:"aura_name"`              // TODO: название ауры
	Retirement          string        `json:"retirement"`             // Количество сбережений
	Au                  int           `json:"au"`                     // TODO: ?
	UPhr                bool          `json:"u_phr"`                  // TODO: ?
	ArenaWon            int           `json:"arena_won"`              // Побед на арене
	ArenaLost           int           `json:"arena_lost"`             // Поражений на арене
	InventorySerial     int           `json:"inventory_serial"`       // Номер инвентаря?
	MaxGp               int           `json:"max_gp"`                 // Максимальное количество праны
	Gold                int           `json:"gold"`                   // Количество золота точное
	GoldWe              string        `json:"gold_we"`                // Количество золото точное текстом
	MonstersKilled      int           `json:"monsters_killed"`        // Убито монстров
	DeathCount          int           `json:"death_count"`            // Смертей
	IsArenaAvailable    bool          `json:"is_arena_available"`     // Доступна ли арена
	DA                  bool          `json:"d_a"`                    // TODO: ?
	SA                  bool          `json:"s_a"`                    // TODO: ?
	RA                  bool          `json:"r_a"`                    // TODO: ?
	IsChfAvailable      bool          `json:"is_chf_available"`       // TODO: ?
	ChfPending          string        `json:"chf_pending"`            // TODO: ?
	AgeStr              string        `json:"age_str"`                // Возраст (времени с регистрации)
	Accumulator         float32       `json:"accumulator"`            // Количество зарядов праны
	QuestsCompleted     int           `json:"quests_completed"`       // Количество завершенных квестов
	Aln                 string        `json:"aln"`                    // Общее описание характера? bad/good
	Dir                 string        `json:"dir"`                    // Направление движения
	MId                 string        `json:"m_id"`                   // TODO: ?
	MIdN                string        `json:"m_id_n"`                 // TODO: ?
	MonsterName         string        `json:"monster_name"`           // Имя монстра
	MonsterProgress     uint16        `json:"monster_progress"`       // Прогресс битвы с монстром
	SProgress           float64       `json:"s_progress"`             // TODO: ?
	ArenaSendAfter      int           `json:"arena_send_after"`       // TODO: ?
	DSendAfter          int           `json:"d_send_after"`           // TODO: ?
	SSendAfter          int           `json:"s_send_after"`           // TODO: ?
	RAfter              int           `json:"r_after"`                // TODO: ?
	RT                  interface{}   `json:"r_t"`                    // TODO: ?
	ChfrAfter           int           `json:"chfr_after"`             // TODO: ?
	IsArenaDisabled     bool          `json:"is_arena_disabled"`      // TODO: ?
	AuraTime            int           `json:"aura_time"`              // Время до окончания ауры
	Q2T                 string        `json:"q2t"`                    // TODO: ?
	Q2                  []interface{} `json:"q2"`                     // TODO: ?
	CTown               string        `json:"c_town"`                 // TODO: ?
	Lte                 string        `json:"lte"`                    // TODO: ?
	Poi                 []int         `json:"poi"`                    // Точки интереса (номера столбов)
	InvM                string        `json:"inv_m"`                  // TODO: ?
	InvitesLeft         int           `json:"invites_left"`           // Осталось инвайтов
	Ggender             string        `json:"ggender"`                // Пол (бога?)
	ArenaGodCmdDisabled bool          `json:"arena_god_cmd_disabled"` // Пульт отключен в режиме арены
}
