// #---------------------------------------#
// #--      @team  : PipelineCN         -- #
// #--      @author: Anthony            -- #
// #--      @time: 2023/08/21 0021      -- #
// #---------------------------------------#

package main

import (
	"Go-blog-server/core"
	"Go-blog-server/global"
	"fmt"
)

func main() {
	// 读取项目配置文件
	core.InitConf()
	fmt.Println(global.Config)
}
