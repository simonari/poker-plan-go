package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUserGroup(r *gin.Engine, db *gorm.DB) {
	g := r.Group("/user")

	rc := NewController(db)

	g.POST("/user/register", rc.RegisterUser)
	g.POST("/user/login", rc.LoginUser)
	g.POST("/user/logout", rc.LogoutUser)
}
