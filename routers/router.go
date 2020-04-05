package routers

import (
	"blogweb_gin/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	loadHtml(router)

	//设置session midddleware
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))

	//router.Use(controllers.SessionMiddleware())

	{

		//注册：
		router.GET("/register", controllers.RegisterGet) //跳转到html用的handler.
		router.POST("/register", controllers.RegisterPost)

		//登录
		router.GET("/login", controllers.LoginGet) //跳转到html用的handler.
		router.POST("/login", controllers.LoginPost)

		//首页
		router.GET("/", controllers.HomeGet)

		//退出
		router.GET("/logout", controllers.ExitGet)

		//路由组
		v1 := router.Group("/article")
		{
			//写文章
			v1.GET("/add", controllers.AddArticleGet)
			v1.POST("/add", controllers.AddArticlePost)
			//显示文章内容
			v1.GET("/show/:id", controllers.ShowArticleGet)

			//更新文章
			v1.GET("/update", controllers.UpdateArticleGet)
			v1.POST("/update", controllers.UpdateArticlePost)

			//删除文章
			v1.GET("/delete", controllers.DeleteArticleGet)

		}
		//显示文章内容
		router.GET("/show/:id", controllers.ShowArticleGet)

		//router.GET("/article/:id",controllers.ShowArticleGet)
		//router.GET("/user/:id", controllers.ShowArticleGet)

		//标签
		router.GET("/tags", controllers.TagsGet)

		//相册
		router.GET("/album", controllers.AlbumGet)

		//文件上传
		router.POST("/upload", controllers.UploadPost)

		//关于我
		router.GET("/aboutme", controllers.AboutMeGet)
	}

	return router

}

/*
加载所有html模板
*/
func loadHtml(router *gin.Engine) {
	//这种方式在linux服务器上找不到该路径.
	//router.LoadHTMLGlob("views/*")
	//绝对路径可以找到, 但显然不推荐.
	router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/blogweb_gin/views/*"))

}
