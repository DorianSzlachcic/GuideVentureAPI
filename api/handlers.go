package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetGames(c *gin.Context) {
	games, err := s.db.GetGames()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &games)
}

func (s *Server) GetGame(c *gin.Context) {
	gameId := c.Param("gameId")
	game, err := s.db.GetGame(gameId)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &game)
}

func (s *Server) GetSteps(c *gin.Context) {
	gameId := c.Param("gameId")
	steps, err := s.db.GetSteps(gameId)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &steps)
}

func (s *Server) GetStep(c *gin.Context) {
	gameId := c.Param("gameId")
	stepIndex := c.Param("stepIndex")
	step, err := s.db.GetStep(gameId, stepIndex)
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &step)
}
