package structs

type Pet struct {
	PetName           string `json:"pet_name"`
	PetClass          string `json:"pet_class"`
	PetLevel          string `json:"pet_level"`
	PetIsDead         bool   `json:"pet_is_dead"`
	PetRename         bool   `json:"pet_rename"`
	PetBirthdayString string `json:"pet_birthday_string"`
}
