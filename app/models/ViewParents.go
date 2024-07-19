package models

const ViewNameParents = "peo_vw_parent"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewParents struct {
	UserId         string `json:"user_id"`
	ParentTypeName string `json:"parent_type_name"`
	BranchId       string `json:"branch_id"`
	BranchName     string `json:"branch_name"`
	FullName       string `json:"full_name"`
	UserPhone      string `json:"user_phone"`
	UserOtp        string `json:"user_otp"`
}

// TableName TmpMstRole's table name
func (*ViewParents) TableName() string {
	return ViewNameParents
}
