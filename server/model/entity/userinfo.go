package entity

type UserInfo struct {
	Id      uint   `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Token   string `json:"token"`
}
