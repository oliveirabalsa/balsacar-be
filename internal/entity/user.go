package entity

type User struct {
	ID       uint   `gorm:"primary_key"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
