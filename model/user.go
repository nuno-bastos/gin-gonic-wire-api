package model

import (
	"time"
)

type User struct {
	Id          string     `gorm:"column:Id;type:char(20);primaryKey"`
	Email       string     `gorm:"column:Email;type:nvarchar(50);not null"`
	FirstName   string     `gorm:"column:FirstName;type:nvarchar(50);not null"`
	LastName    string     `gorm:"column:LastName;type:nvarchar(50);not null"`
	CreatedAt   time.Time  `gorm:"column:CreatedAt;type:datetime2;not null"`
	LastLoginAt time.Time  `gorm:"column:LastLoginAt;type:datetime2;not null"`
	EmployeeId  uint       `gorm:"column:EmployeeId;type:int"`
	Profiles    []*Profile `gorm:"many2many:Security.ProfileUsers;joinForeignKey:UserId;joinReferences:ProfileId"`
}

func (User) TableName() string {
	return "Security.Users"
}
