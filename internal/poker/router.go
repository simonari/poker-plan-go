package poker

import (
	"cmd/poker-backend/internal/room"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	room.AddRoomGroup(router, db)

	return router
}
