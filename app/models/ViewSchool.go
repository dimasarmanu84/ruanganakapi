package models

const ViewNameSchool = "sch_vw_school"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewSchool struct {
	SchoolLogo string `json:"school_logo"`
}

// TableName TmpMstRole's table name
func (*ViewSchool) TableName() string {
	return ViewNameSchool
}
