package infrastructure

import (
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
	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)

	r.Gin.GET("/users/:id", func(c *gin.Context) { usersController.Show(c) })
	// r.Gin.GET("/users/", func(c *gin.Context) { usersController.Get(c) })

	presentationsController := controllers.NewPresentationsController(r.DB)

	r.Gin.GET("/presentetions/:id", func(c *gin.Context) { presentationsController.Get(c) })
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
