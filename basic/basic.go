package basic

import (
	"czm/basic/config"
	"czm/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
