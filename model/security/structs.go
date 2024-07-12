package model_security

import "github.com/nuno-bastos/gin-gonic-wire-api/model"

type TupleProfileUser struct {
	Profile *model.Profile `json:"profile"`
	User    *model.User    `json:"user"`
}

type TupleProcessors1 struct {
	EmployeeId uint
	Level      int
}

type TupleProcessors2 struct {
	ManagerId uint
	Level     int
}

type SecurityMatrixCalculationFilterInput struct {
	EmployeeToManagers map[int][]TupleProcessors2
	EmployeeToReports  map[int][]TupleProcessors1
	LevelToEmployees   map[uint][]uint
}

type AllowOrDenyRule struct {
	UserManagerId   string
	EmployeeId      uint
	AccessLevelRead uint32
}

type ProfileUserAllowOrDenyRules struct {
	ProfileId        string
	UserId           string
	OriginFilter     string
	AllowOrDenyRules []AllowOrDenyRule
}
