package user

type User struct {
	Id             int     `db:"id" json:"id"`
	FullName       string  `db:"fullName" json:"fullName"`
	GenderId       *string `db:"genderId" json:"genderId"`
	Address        string  `db:"address" json:"address"`
	Phone          string  `db:"phone" json:"phone"`
	Email          string  `db:"email" json:"email"`
	Password       string  `db:"password" json:"password"`
	ActivationCode *string `db:"activationCode" json:"activationCode,omitempty"`
	Locale         string  `db:"locale" json:"locale"`
	IsActive       bool    `db:"isActive" json:"isActive"`
	IsAdmin        bool    `db:"isAdmin" json:"isAdmin"`
	CreatedAt      string  `db:"createdAt" json:"createdAt"`
}

type Otp struct {
	Phone string `db:"phone" json:"phone"`
	Code  string `db:"code" json:"code"`
}

// Hide password
func (u *User) Normalize() {
	u.Password = ""
}
