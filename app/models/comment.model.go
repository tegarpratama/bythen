package models

import (
	"app/config"
	"time"
)

type Comment struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	PostID     *uint     `json:"post_id"`
	AuthorName string    `json:"author_name" gorm:"type:varchar(100)"`
	Content    string    `json:"content" gorm:"type:text"`
	CreatedAt  time.Time `json:"created_at"`
	Post       Blog      `json:"-" gorm:"foreignkey:PostID"`
}

type CreateCommentReq struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateComment(comment *Comment) error {
	return config.DB.Create(comment).Error
}

func GetCommentsByPostID(id int) ([]Comment, error) {
	var comment []Comment
	if err := config.DB.Where("post_id = ?", id).Find(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}
