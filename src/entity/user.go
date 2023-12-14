package entity

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Username  string
	Password  string
	CreatedAt int64
	UpdatedAt int64
}
