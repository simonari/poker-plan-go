package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{db}
}

func (ct *Controller) RegisterUser(c *gin.Context) {}

func (ct *Controller) LoginUser(c *gin.Context) {}

func (ct *Controller) LogoutUser(c *gin.Context) {}
