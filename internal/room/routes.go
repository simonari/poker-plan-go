package room

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoomGroup(r *gin.Engine, db *gorm.DB) {
	g := r.Group("/room")
	rc := NewRoomController(db)

	g.GET("/", rc.GetRoomsList)
	g.GET("/:id", rc.GetRoom)

	g.POST("/new", rc.NewRoom)
}
