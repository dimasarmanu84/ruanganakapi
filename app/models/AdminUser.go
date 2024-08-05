package models

import (
	"time"
)

const TableNameAppMstAdminUser = "app_mst_admin_user"

// AppMstAdminUser mapped from table <app_mst_admin_user>
type AppMstAdminUser struct {
	AdminUserID        string    `gorm:"column:admin_user_id;primaryKey" json:"admin_user_id"`
	AdminUserName      string    `gorm:"column:admin_user_name" json:"admin_user_name"`
	AdminUserPassword  string    `gorm:"column:admin_user_password" json:"admin_user_password"`
	AdminUserPhone     string    `gorm:"column:admin_user_phone" json:"admin_user_phone"`
	AdminUserEmail     string    `gorm:"column:admin_user_email" json:"admin_user_email"`
	IsActive           bool      `gorm:"column:is_active" json:"is_active"`
	IsBlocked          bool      `gorm:"column:is_blocked" json:"is_blocked"`
	IsDefaultPassword  bool      `gorm:"column:is_default_password" json:"is_default_password"`
	FailedLoginCount   int32     `gorm:"column:failed_login_count" json:"failed_login_count"`
	AdminUserSession   string    `gorm:"column:admin_user_session" json:"admin_user_session"`
	AdminUserLastLogin time.Time `gorm:"column:admin_user_last_login" json:"admin_user_last_login"`
	AdminPhotoProfile  string    `gorm:"column:admin_photo_profile" json:"admin_photo_profile"`
	BranchID           string    `gorm:"column:branch_id" json:"branch_id"`
	PasswordUser       string    `gorm:"column:password_user" json:"password_user"`
}

// TableName AppMstAdminUser's table name
func (*AppMstAdminUser) TableName() string {
	return TableNameAppMstAdminUser
}
