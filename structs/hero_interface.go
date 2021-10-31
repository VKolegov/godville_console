package structs

type Hero interface {
	GetHealth() int
	GetMaxHealth() int
	GetInvNum() int
	GetMaxInvNum() int
	GetPillar() int
	GetTown() string
	GetGold() int
	GetGoldApprox() string

	GetGodName() string
	GetGodPower() int
	GetGodPowerCharges() int
	GetSavings() string
	GetBricks() int
	GetWood() int
	GetClan() string
	GetClanPosition() string
}
