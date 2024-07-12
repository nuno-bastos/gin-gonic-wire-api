package model

type ProfileUser struct {
	UserID    string `gorm:"column:UserId"`
	ProfileID string `gorm:"column:ProfileId"`
}

func (ProfileUser) TableName() string {
	return "Security.ProfileUsers"
}
