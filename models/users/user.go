package users

import "time"

type User struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ID          string `gorm:"column:ID; primary_key; type:char(36);"`
	Password    string `gorm:"column:password; not null; "`
	Name        string `gorm:"column:name; not null; type:varchar(50)"`
	Email       string `gorm:"column:email; not null; type:varchar(100)"`
	Role        string `gorm :"column:role; not null; type:varchar(20)"`
	PhoneNumber string `gorm:"phone_number;varchar(15)"`
}
