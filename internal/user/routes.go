package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUserGroup(r *gin.Engine, db *gorm.DB) {
	g := r.Group("/user")

	rc := NewController(db)

	g.POST("/register", rc.RegisterUser)
	g.POST("/login", rc.LoginUser)
	g.POST("/logout", rc.LogoutUser)
}
