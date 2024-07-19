package models

const ViewNameMilestoneGroup = "mile_vw_milestone_group"

// TmpMstRole mapped from table <tmp_mst_role>
type ViewMilestoneGroup struct {
	ChildId           string `json:"child_id"`
	ChildFullName     string `json:"child_full_name"`
	ItemName          string `json:"item_name"`
	MilestoneNotes    string `json:"milestone_notes"`
	MilestoneTypeName string `json:"milestone_type_name"`
	ScoreName         string `json:"score_name"`
	UrlImage          string `json:"url_image"`
	ReportDate        string `json:"report_date"`
}

// TableName TmpMstRole's table name
func (*ViewMilestoneGroup) TableName() string {
	return ViewNameMilestoneGroup
}
