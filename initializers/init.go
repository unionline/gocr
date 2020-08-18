package initializers

import (
	"gocr/initializers/config"
	"gocr/initializers/db"
	"gocr/initializers/model"
)

// Initialize ...
func Initialize() {
	config.Init()
	db.Init()
	model.Init()
}
