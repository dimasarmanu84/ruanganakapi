package models

const ViewNameMilestone = "mile_vw_milestone"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewMilestone struct {
	ChildId        string `json:"child_id"`
	ChildFullName  string `json:"child_full_name"`
	ItemName       string `json:"item_name"`
	MilestoneNotes string `json:"milestone_notes"`
	ScoreName      string `json:"score_name"`
	UrlImage       string `json:"url_image"`
	ReportDate     string `json:"report_date"`
}

// TableName TmpMstRole's table name
func (*ViewMilestone) TableName() string {
	return ViewNameMilestone
}
