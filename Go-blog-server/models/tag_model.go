package models

type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`           //标签的名称
	Articlex []ArticleModel `gorm:"many2many:article_tag" json:"-"` //关联该标签的文章列表
	// TagType  int            `gorm:"size:1;default:1" json:"tag_type"` //标签的类型
	//Sites    []SiteModel    `gorm:"many2many:site_tag" json:"-"`      // 关联该标签的文章列表

}
