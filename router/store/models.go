package store

type TStore struct {
	ID       int64  `gorm:"column:user_id;primary_key"`
	Name     string `gorm:"column:name"`
	Raiting  int    `gorm:"column:raiting"`
	Logotype string `gorm:"column:logotype"`
	Location string `gorm:"column:location"`
	Desc     string `gorm:"column:description"`
	Banner   string `gorm:"column:banner"`
}

func (TStore) TableName() string {
	return "tstore"
}

type CountOfReturn struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
