//author: richard
package tag

import "mms/src/component/language"

const (
	Schema = "tags"

	StatusActive  = 0x601
	StatusDeleted = 0x602
	StatusInvalid = 0x699
)

type ITag interface {
	CreateTag(tag *Tag) error
	UpdateTag(tag *Tag) error
	QueryTag(tagId string) (*Tag, error)
	QueryTags(where map[string]interface{}, page int, perPage int) ([]Tag, error)
}

type Tag struct {
	Id 	 string 	`json:"tagId"`
	Status int 		`json:"tagStatus"`
	CreateTime int64 `json:"tagCreateTime"`
	UpdateTime int64 `json:"tagUpdateTime"`
	DeleteTime int64 `json:"tagDeleteTime"`
	Name language.Languages `json:"tagName"`
}
