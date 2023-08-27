package models

import (
	"Go-blog-server/models/ctype"
)

type UserModel struct {
	MODEL
	NickName       string           `gorm:"size:42" json:"nock_name"`
	UserName       string           `gorm:"size:36" json:"user_name"`
	Password       string           `gorm:"size:128" json:"password"`
	Avatar         string           `gorm:"size:256" json:"avatar_id"`
	Email          string           `gorm:"size:128" json:"email"`
	Tel            string           `gorm:"size:18" json:"tel"`
	Addr           string           `gorm:"size:64" json:"addr`
	Token          string           `gorm:"size:64" json:"token"`
	IP             string           `gorm:"size:20" json:"ip"`
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6) json:"sign_status"`
	ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"-"`
	CollectsModels []ArticleModel   `gorm:"many2ymany:auth2_collects;joinForeignKey:AuthID;JoinReferences:ArticleID" json:"-"`
}
