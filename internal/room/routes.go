package room

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoomGroup(r *gin.Engine, db *gorm.DB) {
	g := r.Group("/room")

	rc := NewController(db)

	g.GET("/", rc.GetRoomsList)

	g.POST("/new", rc.NewRoom)

	g.GET("/:id", rc.GetRoom)
	g.DELETE("/:id", rc.DeleteRoom)
	g.PATCH("/:id", rc.PatchRoom)
}
