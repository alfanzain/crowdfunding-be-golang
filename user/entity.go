package user

import "time"

type User struct {
	Id             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFilename string
	Role           string
	CreatedAt      time.Time
	CreatedBy      string
	UpdatedAt      time.Time
	UpdatedBy      string
}
