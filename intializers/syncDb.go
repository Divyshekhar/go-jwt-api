package intializers

import (
	"github.com/Divyshekhar/go-jwt-api/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
