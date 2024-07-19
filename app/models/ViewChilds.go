package models

const ViewChilds = "peo_vw_child"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewChild struct {
	ChildId       string `json:"child_id"`
	ChildFullName string `json:"child_full_name"`
	ItemName      string `json:"item_name"`
	BranchId      string `json:"branch_id"`
	BranchName    string `json:"branch_name"`
	ChildDOB      string `json:"child_dob"`
	Age           string `json:"age"`
}

// TableName TmpMstRole's table name
func (*ViewChild) TableName() string {
	return ViewChilds
}
