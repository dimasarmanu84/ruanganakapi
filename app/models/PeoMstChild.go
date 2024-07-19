package models

import (
	"time"
)

const TableNamePeoMstChild = "peo_mst_child"

// PeoMstChild mapped from table <peo_mst_child>
type PeoMstChild struct {
	ChildID       string    `gorm:"column:child_id;primaryKey" json:"child_id"`
	ChildFullName string    `gorm:"column:child_full_name" json:"child_full_name"`
	ChildNickName string    `gorm:"column:child_nick_name" json:"child_nick_name"`
	ChildDob      time.Time `gorm:"column:child_dob" json:"child_dob"`
	FileID        string    `gorm:"column:file_id" json:"file_id"`
	BranchID      string    `gorm:"column:branch_id" json:"branch_id"`
	ParentID      string    `gorm:"column:parent_id" json:"parent_id"`
	UserID        string    `gorm:"column:user_id" json:"user_id"`
}

// TableName PeoMstChild's table name
func (*PeoMstChild) TableName() string {
	return TableNamePeoMstChild
}
