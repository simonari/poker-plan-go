package room

import (
	"cmd/poker-backend/internal/database"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (ct *Controller) GetRoomsList(c *gin.Context) {
	rooms := new([]database.Room)

	ct.db.Find(&rooms)

	c.JSON(http.StatusOK, rooms)
}

func (ct *Controller) GetRoom(c *gin.Context) {
	roomID := c.Param("id")

	room := new(database.Room)

	result := ct.db.First(&room, roomID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room was not found"})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (ct *Controller) NewRoom(c *gin.Context) {
	requestBody := new(struct {
		Name  string
		Scale int
	})

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	obj := database.Room{
		Name:  requestBody.Name,
		Url:   generateRandomUrl(12),
		Scale: requestBody.Scale,
	}

	result := ct.db.Create(&obj)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, obj)
}

func generateRandomUrl(length int) string {
	//goland:noinspection SpellCheckingInspection
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const numbers = "0123456789"

	result := make([]byte, length)

	for i := range result {
		c := rand.Int() % 2

		fmt.Println(c)

		if c == 1 {
			result[i] = letters[rand.Intn(len(letters))]
		} else {
			result[i] = numbers[rand.Intn(len(numbers))]
		}
	}

	return string(result)
}

func (ct *Controller) DeleteRoom(c *gin.Context) {}

func (ct *Controller) PatchRoom(c *gin.Context) {}
