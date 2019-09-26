package basic

import (
	"github.com/czm448281588/go-micro-test/basic/config"
	"github.com/czm448281588/go-micro-test/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
