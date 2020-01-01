//author: richard
package tag

const (
	Schema = "tags"

	StatusActived = 0x601
	StatusDeleted = 0x602
)

type ITag interface {
	CreateTag(*Tag) error
	DeleteTag(...*Tag) error
	UpdateTag(*Tag) error
	QueryTag(string) (*Tag, error)
}


type Tag struct {
	TagId 	string 	`json:"tagId"`
	TagName string 	`json:"tagName"`
	TagStatus int 	`json:"tagStatus"`
	TagCreateTime int64 `json:"tagCreateTime"`
	TagUpdateTime int64 `json:"tagUpdateTime"`
	TagDeleteTime int64 `json:"tagDeleteTime"`
}
