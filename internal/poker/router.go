package poker

import (
	"cmd/poker-backend/internal/room"
	"cmd/poker-backend/internal/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	room.AddRoomGroup(router, db)
	user.AddUserGroup(router, db)

	return router
}
