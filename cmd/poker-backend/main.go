package main

import (
	"cmd/poker-backend/internal/config"
	"cmd/poker-backend/internal/database"
	"cmd/poker-backend/internal/poker"

	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("%e", err)
	}

	cfg := config.NewConfig()

	db := database.NewConnection(cfg.Database)
	migrateDatabaseAtStartup(db)

	router := poker.SetUpRouter(db)
	router.Run()
}

func migrateDatabaseAtStartup(db *gorm.DB) {
	db.AutoMigrate(
		&database.User{},
		&database.Room{},
		&database.RoomUserRef{},
		&database.Task{},
		&database.RoomTasksRef{},
		&database.Estimate{},
		&database.TaskEstimatesRef{},
	)
}
