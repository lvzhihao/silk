package models

import (
	"github.com/jinzhu/gorm"
)

//go:generate goqueryset -in accounts.go
// 帐号列表
// gen:qs
type Account struct {
	gorm.Model
	Platform  string `gorm:"size:20;unique_index:idx_platform_accountid_serialno" json:"platform"`    //平台
	AccountId string `gorm:"size:100;unique_index:idx_platform_accountid_serialno" json:"account_id"` //平台ID
	SerialNo  string `gorm:"size:100;unique_index:idx_platform_accountid_serialno" json:"serial_no"`  //设备序列号
	NickName  string `gorm:"size:200" json:"nick_name"`                                               //设备昵称
	HeadImage string `gorm:"size:500" json:"head_image"`                                              //设备头像
	QrCode    string `gorm:"size:500" json:"qr_code"`                                                 //帐号二维码
	Status    int32  `gorm:"index:idx_status" json:"status"`                                          //设备状态
	Metadata  string `sql:"type:text" json:"meta_data"`                                               //元数据
}
