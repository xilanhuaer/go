package entity

type EditPassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
