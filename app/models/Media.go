package models

const TableNameTrxMedia = "file_trx_media"

// SchMstBranch mapped from table <sch_mst_branch>
type Media struct {
	FileId       string `gorm:"column:file_id" json:"file_id"`
	FileFullPath string `gorm:"column:file_full_path" json:"file_full_path"`
	FileType     string `gorm:"column:file_type" json:"file_type"`
}

// TableName SchMstBranch's table name
func (*Media) TableName() string {
	return TableNameTrxMedia
}
