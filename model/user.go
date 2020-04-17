package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckUser(username string, password string) bool {
	return true
}


