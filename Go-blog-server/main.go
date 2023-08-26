// #---------------------------------------#
// #--      @team  : PipelineCN         -- #
// #--      @author: Anthony            -- #
// #--      @time: 2023/08/21 0021      -- #
// #---------------------------------------#

package main

import (
	"Go-blog-server/core"
	"Go-blog-server/global"
	"Go-blog-server/router"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	// 读取项目配置文件
	core.InitConf()
	fmt.Println(global.Config)

	// 初始化日志.
	global.Log = core.InitLogger()
	global.Log.Warnln("This if warning.")
	global.Log.Error("This if error.")
	global.Log.Info("This if info.")

	logrus.Warnln("This if warning.")
	logrus.Error("This if error.")
	logrus.Info("This if info.")

	// 连接数据库.
	global.DB = core.InitGorm()
	fmt.Println(global.DB)

	router := router.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_server 运行在: %s", addr)
	router.Run(addr)
}
