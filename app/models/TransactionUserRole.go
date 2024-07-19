package models

const TableNameAppTrxUserRole = "app_trx_user_role"

// AppTrxUserRole mapped from table <app_trx_user_role>
type TransactionUserRole struct {
	UserRoleID  string `gorm:"column:user_role_id;primaryKey" json:"user_role_id"`
	UserID      string `gorm:"column:user_id" json:"user_id"`
	RoleID      string `gorm:"column:role_id" json:"role_id"`
	AdminUserID string `gorm:"column:admin_user_id" json:"admin_user_id"`
}

// TableName AppTrxUserRole's table name
func (*TransactionUserRole) TableName() string {
	return TableNameAppTrxUserRole
}
