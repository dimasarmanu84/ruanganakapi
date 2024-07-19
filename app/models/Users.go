// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameAppMstUser = "app_mst_user"

// AppMstUser mapped from table <app_mst_user>
type AppMstUser struct {
	UserID       string    `gorm:"column:user_id;primaryKey" json:"user_id"`
	UserName     string    `gorm:"column:user_name" json:"user_name"`
	UserEmail    string    `gorm:"column:user_email" json:"user_email"`
	UserPassword string    `gorm:"column:user_password" json:"user_password"`
	UserOtp      int64     `gorm:"column:user_otp" json:"user_otp"`
	IsBlocked    bool      `gorm:"column:is_blocked" json:"is_blocked"`
	IsVerify     bool      `gorm:"column:is_verify" json:"is_verify"`
	DateJoined   time.Time `gorm:"column:date_joined" json:"date_joined"`
	DateExpired  time.Time `gorm:"column:date_expired" json:"date_expired"`
	DateLogin    time.Time `gorm:"column:date_login" json:"date_login"`
	UserIP       string    `gorm:"column:user_ip" json:"user_ip"`
	UserDevice   *string    `gorm:"column:user_device;default:null"  json:"user_device"`
	UserSession  *string    `gorm:"type:string; default:null" json:"user_session"`
	UserPhone    string    `gorm:"column:user_phone" json:"user_phone"`
}

// TableName AppMstUser's table name
func (*AppMstUser) TableName() string {
	return TableNameAppMstUser
}
