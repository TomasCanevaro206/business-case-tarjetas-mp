package routes

import (
	"database/sql"

	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/cmd/api/handler"
	"github.com/TomasCanevaro206/business-case-tarjetas-mp.git/internal/card"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildCardRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildCardRoutes() {

	repo := card.NewRepository(r.db)
	service := card.NewService(repo)
	handler := handler.NewCard(service)
	group := r.rg.Group("/cards")

	group.GET("/", handler.GetAll())
	group.GET("/:id", handler.Get())
	group.POST("/", handler.Create())
	group.PATCH("/:id", handler.Update())
	group.DELETE("/:id", handler.Delete())

}
