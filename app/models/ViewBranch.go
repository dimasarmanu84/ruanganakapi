package models

const ViewNameBranch = "sch_vw_branch"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewBranch struct {
	BranchID   string `json:"branch_id"`
	BranchName string `json:"branch_name"`
	SchoolLogo string `json:"school_logo"`
	SchoolId   string `json:"school_id"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

// TableName TmpMstRole's table name
func (*ViewBranch) TableName() string {
	return ViewNameBranch
}
