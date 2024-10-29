package user

import (
	"cmd/poker-backend/internal/database"
	jwttoken "cmd/poker-backend/internal/utils/token/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

type credentialsType int

const (
	isUsernameCredentials credentialsType = iota
	isEmailCredentials
)

func (ct *Controller) RegisterUser(c *gin.Context) {
	requestBody := new(struct {
		Username     string
		Email        string
		PasswordHash string
	})

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
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

func (ct *Controller) LoginUser(c *gin.Context) {
	requestBody := new(struct {
		UsernameOrEmail string
		PasswordHash    string
	})

	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user := ct.getUserWithCredentials(requestBody.UsernameOrEmail)

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	if user.PasswordHash != requestBody.PasswordHash {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	jwtToken, err := jwttoken.GenerateToken(user.ID)

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt-token": jwtToken})
}

func (ct *Controller) getUserWithCredentials(cred string) *database.User {
	var user *database.User

	if defineUsernameOrEmail(cred) == isUsernameCredentials && validateUsername(cred) {
		user = ct.getUserWithUsername(cred)
	} else if defineUsernameOrEmail(cred) == isEmailCredentials && validateEmail(cred) {
		user = ct.getUserWithEmail(cred)
	}

	return user
}

func defineUsernameOrEmail(s string) credentialsType {
	if strings.Contains(s, "@") {
		return isEmailCredentials
	} else {
		return isUsernameCredentials
	}
}

func (ct *Controller) getUserWithEmail(email string) *database.User {
	user := new(database.User)

	result := ct.db.First(&user, "email = ?", email)

	if result.Error != nil {
		return nil
	}

	return user
}

func (ct *Controller) getUserWithUsername(username string) *database.User {
	user := new(database.User)

	result := ct.db.First(&user, "username = ?", username)

	if result.Error != nil {
		return nil
	}

	return user
}

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
