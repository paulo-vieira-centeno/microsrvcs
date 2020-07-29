package domain

type User struct {
	Id    int64  `json:"id"`
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Email string `json:"email"`
}
