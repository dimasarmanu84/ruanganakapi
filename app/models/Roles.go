package models

const TableNameRoles = "dss_main.tmp_mst_role"

// TmpMstRole mapped from table <tmp_mst_role>
type Roles struct {
	RoleID      uint   `gorm:"column:role_id;primaryKey;autoIncrement:true" json:"role_id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	IsDeleted   string `gorm:"column:is_deleted;default:N" json:"is_deleted"`
	IsActive    string `gorm:"column:is_active;default:Y" json:"is_active"`
	TimeStamps
}

type DataTables struct {
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	Search  string `json:"search"`
	Search2 string `json:"search2"`
}

// TableName TmpMstRole's table name
func (*Roles) TableName() string {
	return TableNameRoles
}
