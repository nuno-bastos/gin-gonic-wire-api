package model

type Employee struct {
	Id             uint   `gorm:"column:Id;primaryKey"`
	EmployeeNumber string `gorm:"column:EmployeeNumber;type:nvarchar(50); not null"`
	FullName       string `gorm:"column:FullName;type:nvarchar(100)"`
	PostTitle      string `gorm:"column:PostTitle;type:nvarchar(40)"`
	Location       string `gorm:"column:Location;type:nvarchar(40)"`
	IsManager      bool   `gorm:"column:IsManager"`
}

func (Employee) TableName() string {
	return "Core.Employees"
}
