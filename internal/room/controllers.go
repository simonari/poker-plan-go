package room

import (
	"cmd/poker-backend/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type RoomController struct {
	db *gorm.DB
}

func NewRoomController(db *gorm.DB) *RoomController {
	return &RoomController{db}
}

func (rc *RoomController) GetRoomsList(c *gin.Context) {
	rooms := new([]database.Room)

	rc.db.Find(&rooms)

	c.JSON(http.StatusOK, rooms)
}

func (rc *RoomController) GetRoom(c *gin.Context) {
	roomID := c.Param("id")

	room := new(database.Room)

	result := rc.db.First(&room, roomID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room was not found"})
		return
	}

	c.JSON(http.StatusOK, room)
}

type NewRoomHolder struct {
	Name  string
	Scale int
}

func (rc *RoomController) NewRoom(c *gin.Context) {
	requestBody := new(NewRoomHolder)

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	url := slug.Make(requestBody.Name)

	obj := database.Room{
		Name:  requestBody.Name,
		Url:   url,
		Scale: requestBody.Scale,
	}

	result := rc.db.Create(&obj)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, obj)
}
