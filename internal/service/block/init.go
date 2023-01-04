package block

import (
	"github.com/fat-garage/wordblock-backend/internal/service"
	"gorm.io/gorm"
)

var Srv service.Block

func Init(db *gorm.DB) {
	Srv = newSrv(db)
}
