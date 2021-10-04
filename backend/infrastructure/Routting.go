package infrastructure

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/DaichiHoshina/repgram_gin/backend/interfaces/controllers"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()

	// g := gin.Default()

	// g.Use(cors.Default())

	// // Corsの設定
	// g.Use(cors.New(cors.Config{
	// 	// アクセスを許可したいアクセス元
	// 	AllowOrigins: []string{
	// 		"http://localhost:3000",
	// 		"http://localhost:3002",
	// 		"https://repgram.com",
	// 	},
	// 	// アクセスを許可したいHTTPメソッド
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"PUT",
	// 		"DELETE",
	// 		"PATCH",
	// 		"OPTIONS",
	// 	},
	// 	// 許可したいHTTPリクエストヘッダ
	// 	AllowHeaders: []string{
	// 		"Access-Control-Allow-Credentials",
	// 		"Access-Control-Allow-Headers",
	// 		"Content-Type",
	// 		"Content-Length",
	// 		"Accept-Encoding",
	// 		"Authorization",
	// 		"X-CSRF-Token",
	// 	},
	// 	// cookieなどの情報を必要とするかどうか
	// 	AllowCredentials: true,
	// }))

	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}

	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	// CORS設定
	r.Gin.Use(cors.Default())

	// ユーザー
	usersController := controllers.NewUsersController(r.DB)
	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Show(c) })
	r.Gin.GET("/auth/user", func(c *gin.Context) { usersController.Connect(c) })
	r.Gin.POST("/auth/login", func(c *gin.Context) { usersController.Login(c) })
	r.Gin.GET("/auth/logout", func(c *gin.Context) { usersController.Logout(c) })
	r.Gin.POST("/auth", func(c *gin.Context) { usersController.Create(c) })

	// 投稿
	presentationsController := controllers.NewPresentationsController(r.DB)
	r.Gin.GET("/presentations/:id", func(c *gin.Context) { presentationsController.Show(c) })
	r.Gin.GET("/presentations", func(c *gin.Context) { presentationsController.Index(c) })
	r.Gin.POST("/presentations", func(c *gin.Context) { presentationsController.Create(c) })
	r.Gin.PUT("/presentations/:id", func(c *gin.Context) { presentationsController.Update(c) })
	r.Gin.DELETE("/presentations/:id", func(c *gin.Context) { presentationsController.Delete(c) })

	// いいね機能
	likesController := controllers.NewLikesController(r.DB)
	r.Gin.POST("/likes", func(c *gin.Context) { likesController.Create(c) })
	r.Gin.POST("/likes/delete", func(c *gin.Context) { likesController.Delete(c) })

	// ヘルスチェック
	r.Gin.GET("/health_check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
