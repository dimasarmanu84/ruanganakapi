package models

const ViewNameTimesheetEducator = "time_vw_educator_timesheet"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewEducatorTimesheet struct {
	UsetId     string `json:"user_id"`
	Username   string `json:"user_name"`
	BranchName string `json:"branch_name"`
	ClockIn    string `json:"clock_in"`
	ClockOut   string `json:"clock_out"`
}

// TableName TmpMstRole's table name
func (*ViewEducatorTimesheet) TableName() string {
	return ViewNameTimesheetEducator
}
