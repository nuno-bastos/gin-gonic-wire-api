package model

type Profile struct {
	Id              string  `gorm:"column:Id;type:char(20);primaryKey"`
	ProfileTypeId   string  `gorm:"column:ProfileTypeId;type:char(20)"`
	Name            string  `gorm:"column:Name;type:nvarchar(30);not null"`
	Description     string  `gorm:"column:Description;type:nvarchar(200);not null"`
	Level           uint    `gorm:"column:Level;not null"`
	Division        uint    `gorm:"column:Division"`
	AccessLevelRead uint    `gorm:"column:AccessLevelRead;not null"`
	Users           []*User `gorm:"many2many:Security.ProfileUsers;joinForeignKey:ProfileId;joinReferences:UserId"`
}

func (Profile) TableName() string {
	return "Security.Profiles"
}
