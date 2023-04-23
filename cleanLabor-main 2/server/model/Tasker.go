package model

type Tasker struct {
	UserName string `json: "UserName"`
	Desc string `json: "desc"`
	WorkContent string `json: "WorkContent"`
	Phone string `json: "Phone"`
	Email    string `json:"email"`
    Password string `json:"password"`
}

