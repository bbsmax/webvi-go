package user

import "time"

//User 데이터베이스 모델
type User struct {
	CreatedAt   time.Time `gorm:"primary_key; default:0"`
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ID          string `gorm:"column:ID; primary_key; type:char(36);"`
	Password    string `gorm:"column:password; not null; "`
	Name        string `gorm:"column:name; not null; type:varchar(50)"`
	Email       string `gorm:"column:email; not null; type:varchar(100)"`
	Role        string `gorm:"column:role; not null; type:varchar(20)"`
	PhoneNumber string `gorm:"column:phone_number; type:varchar(15)"`
	Avator      string `gorm:"column:avator; type:varchar(100);"`
}

func (User) TableName() string {
	return "user"
}
