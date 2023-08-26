package settings_api

import (
	"Go-blog-server/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.Ok(map[string]string{}, "Hello World!", c)

	res.OkWithData(map[string]string{
		"id":       "0000-1111-2222-3333",
		"language": "zh_CN",
		"env":      "GoLang21.2",
	}, c)

	res.FailWithCode(res.SettingsError, c)
	// c.JSON(200, gin.H{"msg": "Hello GoLang!"})
}
