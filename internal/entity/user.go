package entity

type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
