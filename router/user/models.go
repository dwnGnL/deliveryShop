package user

type User struct {
	ID          int64  `gorm:"column:user_id;primary_key"`
	Username    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	FullName    string `gorm:"column:fullname"`
	PhoneNumber string `gorm:"column:phone"`
	Salt        string `gorm:"column:salt"`
	Role        int    `gorm:"column:role"`
	Disabled    bool   `gorm:"column:disabled"`
}

func (User) TableName() string {
	return "tusers"
}

type CreateUser struct {
	UserName    string `json:"userName"`
	Password    string `json:"password"`
	FIO         string `json:"fio"`
	PhoneNumber string `json:"phone"`
}
