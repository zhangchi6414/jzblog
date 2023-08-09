package model

type Tag struct {
	*Model
	Name  string `json:"name;comment:标签名'"`
	State uint   `json:"state;comment:状态,0为禁用，1为启动;default:0"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
