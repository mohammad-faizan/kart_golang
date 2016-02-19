package data

import(
	"time"
)

type User struct {
	Id 						int 				`db:"id" 						json:"id"`
	FirstName 		string 			`db:"first_name" 		json:"first_name"`
	LastName 			string 			`db:"last_name" 		json:"last_name"`
	Dob 					time.Time 	`db:"dob" 					json:"dob"`
	Gender 				string 			`db:"gender" 				json:"gender"`
	Email 				string 			`db:"email" 				json:"email"`

	password 			string
	passwordConfirmation string
}

func (u *User) SetPassword(p string) {
	u.password = p
}

func (u *User) SetPasswordConfirmation(p string) {
	u.passwordConfirmation = p
}

func (u *User) PasswordHash() string {
	return u.password
}