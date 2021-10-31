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
}
