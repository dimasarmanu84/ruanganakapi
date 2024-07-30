package models

const ViewNameEducatorBranch = "sch_vw_educator_branch"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewEducatorBranch struct {
	EducatorBranchId string `json:"educator_branch_id"`
	EducatorId       string `json:"educator_id"`
	BranchId         string `json:"user_name"`
	BranchName       string `json:"branch_name"`
	UserId           string `json:"user_id"`
}

// TableName TmpMstRole's table name
func (*ViewEducatorBranch) TableName() string {
	return ViewNameEducatorBranch
}
