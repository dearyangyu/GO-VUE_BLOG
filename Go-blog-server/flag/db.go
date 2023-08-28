package flag

import (
	"Go-blog-server/global"
	"Go-blog-server/models"
)

func Makemigrations() {
	// var err error

	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.User2Collects{})
	// global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	// // 生成四张表的表结构
	// err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
	// 	AutoMigrate(
	// 		&models.BannerModel{},
	// 		&models.TagModel{},
	// 		&models.MessageModel{},
	// 		&models.AdvertModel{},
	// 		&models.CommentModel{},
	// 		&models.ArticleModel{},
	// 		&models.UserModel{},
	// 		&models.MenuModel{},
	// 		&models.MenuBannerModel{},
	// 		&models.FadeBackModel{},
	// 		&models.LoginDataModel{},
	// 	)

	// if err != nil {
	// 	global.Log.Error("[ERROR]: 生成数据表结构失败!")
	// 	return
	// }
	// global.Log.Info("[SUCCESS]: 生成数据库表结构成功!")
}
