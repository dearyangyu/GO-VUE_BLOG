package router

import (
	"Go-blog-server/api"
)

func (router RouterGroup) SettingsRouter() {

	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsInfoView)

}
