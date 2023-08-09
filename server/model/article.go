package model

type Article struct {
	*Model
	Title         string `json:"title;comment:文章标题"`
	Desc          string `json:"desc;comment:文章简介"`
	CoverImageUrl string `json:"cover_image_url;comment:图片地址"`
	Content       string `json:"content;comment:文章内容"`
	State         uint   `json:"state;comment:状态,0为禁用，1为启动;default:0;type:longtext"`
}

func (t Article) TableName() string {
	return "blog_article"
}
