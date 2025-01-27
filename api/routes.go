package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type name struct {
}

func InitRoutes(e *gin.Engine, db *gorm.DB) {

	r := InitRepository(
		db,
	)
	s := InitService(
		r,
		16,
	)
	h := InitHandler(
		s,
	)

	e.GET("/sources", h.GetData)
	e.GET("/update", h.QueueUpdate)
}
