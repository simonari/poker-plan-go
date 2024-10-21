package room

import (
	"cmd/poker-backend/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (rc *Controller) GetRoomsList(c *gin.Context) {
	rooms := new([]database.Room)

	rc.db.Find(&rooms)

	c.JSON(http.StatusOK, rooms)
}

func (rc *Controller) GetRoom(c *gin.Context) {
	roomID := c.Param("id")

	room := new(database.Room)

	result := rc.db.First(&room, roomID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room was not found"})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (rc *Controller) NewRoom(c *gin.Context) {
	requestBody := new(struct {
		Name  string
		Scale int
	})

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
