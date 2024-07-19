package models

const TableNameSchTrxStoryLikeComment = "sch_trx_story_like_comment"

// SchTrxStoryLikeComment mapped from table <sch_trx_story_like_comment>
type TransactionStoryLikeComment struct {
	LikeCommentID string `gorm:"column:like_comment_id;primaryKey" json:"like_comment_id"`
	StoryID       string `gorm:"column:story_id" json:"story_id"`
	UserID        string `gorm:"column:user_id" json:"user_id"`
	IsLike        bool   `gorm:"column:is_like" json:"is_like"`
	Comment       string `gorm:"column:comment" json:"comment"`
}

// TableName SchTrxStoryLikeComment's table name
func (*TransactionStoryLikeComment) TableName() string {
	return TableNameSchTrxStoryLikeComment
}
