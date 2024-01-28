package model

type User struct {
	Phone string `json:"phone"`
	Username string `json:"username"`
}

func (u *User) Validate() bool {
	if u.Phone == "" || u.Username == "" {
		return false
	}
	return true 
}