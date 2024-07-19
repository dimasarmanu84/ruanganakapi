package models

const ViewNameTimesheetChild = "time_vw_child_timesheet"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewChildTimesheet struct {
	ChildId       string `json:"child_id"`
	ChildFullName string `json:"child_full_name"`
	BranchName    string `json:"branch_name"`
	SchoolName    string `json:"school_name"`
	ClockIn       string `json:"clock_in"`
	ClockOut      string `json:"clock_out"`
}

// TableName TmpMstRole's table name
func (*ViewChildTimesheet) TableName() string {
	return ViewNameTimesheetChild
}
