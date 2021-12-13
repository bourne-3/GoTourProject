package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

// 专门用于展示Swagger文档
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
