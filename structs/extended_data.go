package structs

import "time"

type ExtendedData struct {
	Status string  `json:"status"`
	Hero   HeroObj `json:"hero"`
	Skills []Skill `json:"skills"`
	Inventory map[string]InventoryItem `json:"inventory"`
	Equipment Equipment                `json:"equipment"`
	Ctime     time.Time                `json:"ctime"`
	Gca       struct {
		Absnt  int `json:"absnt"`
		Absnt2 int `json:"absnt_2"`
		Absnt3 int `json:"absnt_3"`
		Reglr  int `json:"reglr"`
		Tmpl3  int `json:"tmpl_3"`
		Gcard  int `json:"gcard"`
		Dth    int `json:"dth"`
		Train  int `json:"train"`
		Evilm  int `json:"evilm"`
		Sk5    int `json:"sk5"`
		Bmk    int `json:"bmk"`
		Reglr2 int `json:"reglr_2"`
		Sk52   int `json:"sk5_2"`
		Evilm2 int `json:"evilm_2"`
		Tmpl2  int `json:"tmpl_2"`
		Reglr3 int `json:"reglr_3"`
		Bmk2   int `json:"bmk_2"`
		Pet3   int `json:"pet_3"`
		Dth2   int `json:"dth_2"`
		Goodm2 int `json:"goodm_2"`
		Tmpl   int `json:"tmpl"`
		Absnt4 int `json:"absnt_4"`
		Rtrm   int `json:"rtrm"`
		Reglr4 int `json:"reglr_4"`
		Q2     int `json:"q2"`
		Q22    int `json:"q2_2"`
		Evilm3 int `json:"evilm_3"`
		Wood3  int `json:"wood_3"`
		Wood2  int `json:"wood_2"`
	} `json:"gca"` // Прогресс по заслугам
	Gcak struct {
		Reglr  string `json:"reglr"`
		Sk5    string `json:"sk5"`
		Evilm  string `json:"evilm"`
		Tmpl3  string `json:"tmpl_3"`
		Reglr2 string `json:"reglr_2"`
		Bmk    string `json:"bmk"`
		Dth    string `json:"dth"`
		Pet3   string `json:"pet_3"`
		Tmpl2  string `json:"tmpl_2"`
		Absnt4 string `json:"absnt_4"`
		Absnt  string `json:"absnt"`
		Absnt2 string `json:"absnt_2"`
		Absnt3 string `json:"absnt_3"`
		Evilm2 string `json:"evilm_2"`
		Tmpl   string `json:"tmpl"`
		Q2     string `json:"q2"`
		Reglr3 string `json:"reglr_3"`
		Gcard  string `json:"gcard"`
		Wood3  string `json:"wood_3"`
	} `json:"gcak"`
	Gcid          string        `json:"gcid"` // TODO: ?
	Pet           Pet           `json:"pet"`
	HasPet        bool          `json:"has_pet"`
	Hints         []interface{} `json:"hints"` // TODO: ?
	NewsFromField struct {
		Time time.Time `json:"time"`
		Msg  string    `json:"msg"`
	} `json:"news_from_field"` // Вести с полей
	Diary []struct {
		Time time.Time `json:"time"`
		Msg  string    `json:"msg"`
		I    int       `json:"i"`
		Infl bool      `json:"infl,omitempty"` // Результат влияния
	} `json:"diary"`
	ImpE []struct {
		Time      time.Time `json:"time"`
		Msg       string    `json:"msg"`
		FId       string    `json:"f_id,omitempty"`
		Arena     string    `json:"arena,omitempty"`
		I         int       `json:"i"`
		UphAuthor string    `json:"uph_author,omitempty"`
	} `json:"imp_e"` // Важные записи в дневнике
	GcMid  int    `json:"gc_mid"`   // TODO: ?
	GcMidK int    `json:"gc_mid_k"` // TODO: ?
	GcN    string `json:"gc_n"`     // TODO: ?
	Lbp    int    `json:"lbp"`      // TODO: ?
}
