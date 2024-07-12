package model_security

type GoMatrixRule struct {
	UserId          string `gorm:"column:UserId;primaryKey"`
	EmployeeId      uint   `gorm:"column:EmployeeId;primaryKey"`
	AccessLevelRead uint32 `gorm:"column:AccessLevelRead"`
}
