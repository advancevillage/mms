//author: richard
package tag

const (
	Schema = "tags"

	StatusActived = 0x301
	StatusDeleted = 0x302
)

type ITag interface {
	CreateTag(*Tag) error
	DeleteTag(...*Tag) error
	UpdateTag(*Tag) error
	QueryTag(int64) (*Tag, error)
}


type Tag struct {
	TagId 	int64 	`json:"tagId"`
	TagName string 	`json:"tagName"`
	TagStatus int 	`json:"tagStatus"`
	TagCreateTime int64 `json:"tagCreateTime"`
	TagUpdateTime int64 `json:"tagUpdateTime"`
	TagDeleteTime int64 `json:"tagDeleteTime"`
}
