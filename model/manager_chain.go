package model

type ManagerChain struct {
	Id             uint `gorm:"column:Id;primaryKey"`
	ManagerId      int  `gorm:"column:ManagerId;not null"`
	EmployeeId     int  `gorm:"column:EmployeeId;not null"`
	Level          int  `gorm:"column:Level;not null"`
	EmployeeStatus bool `gorm:"column:EmployeeStatus;not null"`
}

func (ManagerChain) TableName() string {
	return "Core.ManagerChains"
}
