package models

const TableNameMenu = "dss_main.tmp_mst_menu"

type Menu struct {
	MenuID        uint   `gorm:"primary_key" json:"menu_id"`
	Name          string `json:"name"`
	Class         string `json:"class"`
	Link          string `json:"link"`
	Level         uint   `json:"level"`
	Parent        uint   `json:"parent"`
	Icon          string `json:"icon"`
	Counter       uint   `json:"counter"`
	IsActive      string `json:"is_active"`
	DisplayOnTree string `json:"display_on_tree"`
}

// TableName TmpMstFunction's table name
func (*Menu) TableName() string {
	return TableNameMenu
}
