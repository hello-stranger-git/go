package usermodel

import (
	responsemodel "go_project/model/responseModel"
	mysqlutils "go_project/utils/mysqlUtils"

	"github.com/Gre-Z/common/jtime"
	"gorm.io/gorm"
)

type UserInfos struct {
	Id        int             `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt jtime.JsonTime  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt jtime.JsonTime  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
	UserName  string          `gorm:"column:username" json:"username"`
	PassWord  string          `gorm:"column:password" json:"password"`
	SereKey   string          `gorm:"column:secret_key" json:"secret_key"`
	Phone     string          `gorm:"column:phone" json:"phone"`
	Status    int             `gorm:"column:status" json:"status"`
	Remark    string          `gorm:"column:remark" json:"remark"`
}
type DB struct {
	db *gorm.DB
}

// 初始化数据库链接
func GetUserModel() *DB {
	database := mysqlutils.GetDB()
	database.Table("user_infos").Model(&UserInfos{})
	return &DB{db: database}
}

// 注册
func (db *DB) Register(data *UserInfos) (*UserInfos, error) {
	err := db.db.Create(&data).Error
	return data, err
}

// 登陆
func (db *DB) Login(username string) (data *UserInfos, err error) {
	err = db.db.Where("username = ?", username).Find(&data).Error
	return data, err
}

//设置用户ID
func (db *DB) SetId(id uint) *DB {
	db.db = db.db.Where("id = ?", id)
	return db
}

//根据id查询用户
func (db *DB) GetUserInfoById() (data *responsemodel.UserInfo, err error) {
	err = db.db.First(&data).Error
	return data, err
}
