package models

type BannerModel struct {
	MODEL
	Path          string         `json:"path"`                                             //图片路径
	Sort          int            `gorm:"size:4default:0" josn:"sort"`                      //排序字段
	hash          string         `json:"hash"`                                             //图片的hash值，用于判断图片是否重复
	Name          string         `gorm:"size:38" josn:"name"`                              //图片名称
	NotUse        bool           `gorm:"default:false;comment:是否可以选择，默认可以" json:"not_use"` //是否可以选择
	ArticleModels []ArticleModel `gorm:"foreignKey:CoverID" json:"-"`
}
