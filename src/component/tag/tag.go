//author: richard
package tag

import "mms/src/component/language"

const (
	Schema = "tags"

	StatusActived = 0x601
	StatusDeleted = 0x602
)

type ITag interface {
	CreateTag(tag *Tag) error
	DeleteTag(tag ...*Tag) error
	UpdateTag(tag *Tag) error
	QueryTag(tagId string) (*Tag, error)
}


type Tag struct {
	TagId 	string 	`json:"tagId"`
	TagStatus int 	`json:"tagStatus"`
	CreateTime int64 `json:"tagCreateTime"`
	UpdateTime int64 `json:"tagUpdateTime"`
	DeleteTime int64 `json:"tagDeleteTime"`
	TagName language.Languages 	`json:"tagName"`
}
