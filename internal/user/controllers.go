package user

import (
	"cmd/poker-backend/internal/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (ct *Controller) RegisterUser(c *gin.Context) {
	requestBody := new(struct {
		Username     string
		Email        string
		PasswordHash string
	})

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		log.Println(err)
		c.Status(http.StatusBadRequest)
		return
	}

	if !validateUsername(requestBody.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username is invalid"})
		return
	}

	if !validateEmail(requestBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is invalid"})
		return
	}

	result := ct.db.Create(&database.User{
		Username:     requestBody.Username,
		Email:        requestBody.Email,
		PasswordHash: requestBody.PasswordHash,
	})

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (ct *Controller) LoginUser(c *gin.Context) {}

func (ct *Controller) LogoutUser(c *gin.Context) {}

func validateUsername(s string) bool {
	pattern, err := regexp.Compile("^[a-zA-Z0-9]+$")

	if err != nil {
		log.Println(err)
		return false
	}

	return pattern.MatchString(s)
}

func validateEmail(s string) bool {
	pattern, err := regexp.Compile("^\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")

	if err != nil {
		log.Println(err)
		return false
	}

	return pattern.MatchString(s)
}
