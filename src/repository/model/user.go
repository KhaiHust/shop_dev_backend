package model

type User struct {
	BaseModel
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
