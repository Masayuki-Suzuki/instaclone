package types

type ErrorMessage struct {
	ErrorMessage	string	`json:"errorMessage"`
}

type User struct {
	Username	string `json:"username"`
	Uid			string `json:"uid"`
	Email		string `json:"email"`
	FullName	string `json:"fullName"`
}

type SignUpForm struct {
	EmailSignUp bool
	User
}
