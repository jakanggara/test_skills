package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *Service
}

func InitHandler(s *Service) *Handler {
	return &Handler{
		s: s,
	}
}

func (h *Handler) QueueUpdate(c *gin.Context) {

	h.s.Queue()

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Update task has been scheduled",
	})
}

func (h *Handler) GetData(c *gin.Context) {

	s, err := h.s.r.GetAllSource()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": s,
	})
}
