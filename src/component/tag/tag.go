//author: richard
package tag

import "mms/src/language"

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
	QueryTags(where map[string]interface{}, page int, perPage int, sort map[string]interface{}) ([]Tag, int64, error)
}

type Tag struct {
	Id 	 string 	 `json:"id"`
	Status int 		 `json:"status"`
	CreateTime int64 `json:"createTime"`
	UpdateTime int64 `json:"updateTime"`
	DeleteTime int64 `json:"deleteTime"`
	Name *language.Languages `json:"name"`
}
