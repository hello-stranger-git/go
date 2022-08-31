package responsemodel

import (
	"github.com/Gre-Z/common/jtime"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id        int             `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt jtime.JsonTime  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt jtime.JsonTime  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	UserName  string          `gorm:"column:username" json:"username"`
	SereKey   string          `gorm:"column:secret_key" json:"secret_key"`
	Phone     string          `gorm:"column:phone" json:"phone"`
	Status    int             `gorm:"column:status" json:"status"`
	Remark    string          `gorm:"column:remark" json:"remark"`
}
