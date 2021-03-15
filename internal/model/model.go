package model

import (
	"blog-service/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
)

//gong you model
type Model struct{
	ID uint32 `gorm:"primary_key" json:"id"`
	CreateBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS)(*gorm.DB,error){

	db,err := gorm.Open(databaseSetting.DBType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime))
	if err != nil {
		return nil,err
	}
	return db,nil
}
