package responseuserstructs

type LoginResponse struct {
	Token    string      `json:"token"`
	UserInfo interface{} `json:"user_info"`
}
type UserInfo struct {
	Id       int    `json:"id"  gorm:"column:id"`
	Phone    string `json:"phone"  gorm:"column:phone"`
	Username string `json:"username"  gorm:"column:username"`
}
