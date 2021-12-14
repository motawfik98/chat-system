package domain

type Chat struct {
	ID       uint        `json:"-" gorm:"primaryKey"`
	Number   uint        `json:"number" gorm:"autoIncrement;uniqueIndex:number_token"`
	AppToken string      `json:"appToken" gorm:"uniqueIndex:number_token;size:36" validate:"required"`
	App      Application `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AppToken;references:Token" validate:"required,nostructlevel"`
}
