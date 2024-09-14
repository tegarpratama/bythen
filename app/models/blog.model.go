package models

import (
	"app/config"
	"time"
)

type Blog struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar(100)"`
	Content   string    `json:"content" gorm:"type:text"`
	AuthorID  *uint     `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Author    User      `json:"author" gorm:"foreignkey:AuthorID"`
}

type CreateBlogReq struct {
	Title   string `json:"title" inding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateBlog(blog *Blog) error {
	return config.DB.Create(blog).Error
}

func GetAllBlogs(offset int, limit int, search string) ([]Blog, int64, error) {
	var blogs []Blog
	var count int64

	if search != "" {
		if err := config.DB.Where("title LIKE ?", "%"+search+"%").
			Model(&Blog{}).
			Count(&count).
			Error; err != nil {
			return nil, 0, err
		}

		if err := config.DB.Where("title LIKE ?", "%"+search+"%").
			Preload("Author").
			Offset(offset).
			Limit(limit).
			Find(&blogs).
			Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err := config.DB.Model(&Blog{}).Count(&count).Error; err != nil {
			return nil, 0, err
		}

		if err := config.DB.Preload("Author").Offset(offset).Limit(limit).Find(&blogs).Error; err != nil {
			return nil, 0, err
		}
	}

	return blogs, count, nil
}

func GetBlogByID(id int) (Blog, error) {
	var blog Blog
	if err := config.DB.Preload("Author").Where("id = ?", id).First(&blog).Error; err != nil {
		return blog, err
	}

	return blog, nil
}

func UpdateBlog(blog *Blog) error {
	return config.DB.Save(blog).Error
}

func DeleteBlog(id int) error {
	return config.DB.Where("id = ?", id).Delete(&Blog{}).Error
}
