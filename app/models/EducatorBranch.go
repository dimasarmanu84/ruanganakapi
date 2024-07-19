package models

const TableNameSchTrxEducatorBranch = "sch_trx_educator_branch"

// SchTrxEducatorBranch mapped from table <sch_trx_educator_branch>
type SchTrxEducatorBranch struct {
	EducatorBranchID string `gorm:"column:educator_branch_id;primaryKey" json:"educator_branch_id"`

	BranchID   string `gorm:"column:branch_id" json:"branch_id"`
	EducatorID string `gorm:"column:educator_id" json:"educator_id"`
}

// TableName SchTrxEducatorBranch's table name
func (*SchTrxEducatorBranch) TableName() string {
	return TableNameSchTrxEducatorBranch
}
