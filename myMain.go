package main

import (
	"blogweb_gin/database"
	"blogweb_gin/routers"
)

/*
1.补全包的方式:
	新建文件夹:D:\michael.cui\workspace_go\src\golang.org\x\net (注意,不要建到x\net\html这个层. 下面下载后是包括html层的,否则报错:no Go files in...)
	>git clone https://github.com/golang/net D:\michael.cui\workspace_go\src\golang.org\x\net

2.补全Jquery:
	如果没有Jquery的库,那么FE的Jquery会失效. 仔细观察js里的类似'$("register-form").validate({',
	第一个$符号下会有黄线, 光标提示:"Unresolved function or method $()",此时就需要引入Jquery库了.
	方式一(推荐):
		在随意一个html中,写入:<script src="https://code.jquery.com/jquery-3.3.1.min.js"></script>,
			把光标移动到连接上,alt+enter即可自动下载.并且没有墙的问题.
	方式二:
		在idea中选择file-settings--搜索"Libraries"(在在Language&Frameworks/JavaScript/下)--点击右侧的Download(可能需要刷新几次)
			--找到jquery--download and install.

3.页面清缓存:
	ctrl + f5
*/
func main() {
	database.InitMysql()
	router := routers.InitRouter()

	//静态资源
	router.Static("/static", "./static")

	router.Run(":8080")
}
