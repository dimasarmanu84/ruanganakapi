package models

const ViewNameTimesheetChildOvertime = "v_timesheet_child_overtime"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewChildTimesheetOvertime struct {
	ChildId       string `json:"child_id"`
	ChildFullName string `json:"child_full_name"`
	BranchName    string `json:"branch_name"`
	SchoolName    string `json:"school_name"`
	ClockIn       string `json:"clock_in"`
	ClockOut      string `json:"clock_out"`
	TotalOvertime string `json:"total_overtime"`
	OvertimeIn    string `json:"overtimein"`
	OvertimeOut   string `json:"overtimeout"`
}

// TableName TmpMstRole's table name
func (*ViewChildTimesheetOvertime) TableName() string {
	return ViewNameTimesheetChild
}