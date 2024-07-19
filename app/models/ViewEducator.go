package models

const ViewEducatorName = "peo_vw_educator"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewEducator struct {
	EducatorID string `json:"educator_id"`
	FullName   string `json:"full_name"`
	UserID     string `json:"user_id"`
	BranchName string `json:"branch_name"`
	ChildDOB   string `json:"child_dob"`
	UserPhone  string `json:"user_phone"`
	UserOTP    string `json:"user_otp"`
}

// TableName TmpMstRole's table name
func (*ViewEducator) TableName() string {
	return ViewEducatorName
}
